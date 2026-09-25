package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	// ErrRawNotFound indicates the raw source file is missing.
	ErrRawNotFound = errors.New("raw source file not found")
	// ErrTransNotFound indicates the translated HTML file is missing.
	ErrTransNotFound = errors.New("translated html file not found")
)

var (
	reArticleContent = regexp.MustCompile(`(?s)<article[^>]*id="chapter-content"[^>]*>(.*?)</article>`)
	reParagraphTag   = regexp.MustCompile(`(?s)<p>(.*?)</p>`)
	reHTMLTags       = regexp.MustCompile(`<[^>]+>`)
)

// AuditStatus represents the audit outcome for a chapter.
type AuditStatus string

const (
	StatusPass    AuditStatus = "PASS"
	StatusWarn    AuditStatus = "WARN"
	StatusFail    AuditStatus = "FAIL"
	StatusPending AuditStatus = "PENDING"
	StatusNoRaw   AuditStatus = "NO_RAW"
)

// ChapterAudit contains the comparison metrics between raw and translated text.
type ChapterAudit struct {
	ChapterNum     int
	RawParas       int
	TransParas     int
	RawChars       int
	EnglishWords   int
	WordCharRatio  float64
	ParaRatio      float64
	EmptyParagraph int
	Status         AuditStatus
	Detail         string
}

// ExtractRawParagraphs extracts non-blank, non-comment lines from raw text.
func ExtractRawParagraphs(rawText string) []string {
	lines := strings.Split(rawText, "\n")
	paragraphs := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
			paragraphs = append(paragraphs, trimmed)
		}
	}
	return paragraphs
}

// ExtractTranslatedParagraphs parses paragraph texts from chapter HTML content.
func ExtractTranslatedParagraphs(htmlContent string) (paras []string, emptyCount int) {
	targetHTML := htmlContent
	if match := reArticleContent.FindStringSubmatch(htmlContent); len(match) > 1 {
		targetHTML = match[1]
	}

	matches := reParagraphTag.FindAllStringSubmatch(targetHTML, -1)
	paras = make([]string, 0, len(matches))
	for _, m := range matches {
		content := strings.TrimSpace(m[1])
		if content == "" {
			emptyCount++
			continue
		}
		cleanText := reHTMLTags.ReplaceAllString(content, " ")
		cleanText = strings.TrimSpace(cleanText)
		if cleanText != "" {
			paras = append(paras, cleanText)
		} else {
			emptyCount++
		}
	}
	return paras, emptyCount
}

// CountWords counts the number of space-delimited words in a slice of paragraphs.
func CountWords(paragraphs []string) int {
	total := 0
	for _, p := range paragraphs {
		total += len(strings.Fields(p))
	}
	return total
}

// ComputeAudit evaluates the fidelity of a translated chapter against raw text.
func ComputeAudit(chapterNum int, rawText, htmlContent string) ChapterAudit {
	rawParas := ExtractRawParagraphs(rawText)
	rawChars := utf8.RuneCountInString(strings.Join(rawParas, ""))

	transParas, emptyParas := ExtractTranslatedParagraphs(htmlContent)
	englishWords := CountWords(transParas)

	audit := ChapterAudit{
		ChapterNum:     chapterNum,
		RawParas:       len(rawParas),
		TransParas:     len(transParas),
		RawChars:       rawChars,
		EnglishWords:   englishWords,
		EmptyParagraph: emptyParas,
	}

	if audit.RawChars > 0 {
		audit.WordCharRatio = float64(audit.EnglishWords) / float64(audit.RawChars)
	}
	if audit.RawParas > 0 {
		audit.ParaRatio = float64(audit.TransParas) / float64(audit.RawParas)
	}

	switch {
	case audit.EmptyParagraph > 0:
		audit.Status = StatusFail
		audit.Detail = fmt.Sprintf("Contains %d empty <p> tags", audit.EmptyParagraph)
	case audit.WordCharRatio < 0.50 || audit.ParaRatio < 0.60:
		audit.Status = StatusFail
		audit.Detail = fmt.Sprintf("Severely condensed (ratio: %.2f, paras: %d/%d)", audit.WordCharRatio, audit.TransParas, audit.RawParas)
	case audit.WordCharRatio < 0.58 || audit.ParaRatio < 0.75:
		audit.Status = StatusWarn
		audit.Detail = fmt.Sprintf("Partially condensed (ratio: %.2f)", audit.WordCharRatio)
	default:
		audit.Status = StatusPass
		audit.Detail = "Full fidelity verified"
	}

	return audit
}

// ValidateChapter audits a single chapter given file paths.
func ValidateChapter(rawPath, htmlPath string, chapterNum int) (ChapterAudit, error) {
	rawBytes, err := os.ReadFile(rawPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ChapterAudit{
				ChapterNum: chapterNum,
				Status:     StatusNoRaw,
				Detail:     "Raw source file missing",
			}, ErrRawNotFound
		}
		return ChapterAudit{ChapterNum: chapterNum}, fmt.Errorf("read raw chapter %d: %w", chapterNum, err)
	}

	htmlBytes, err := os.ReadFile(htmlPath)
	if err != nil {
		if os.IsNotExist(err) {
			rawParas := ExtractRawParagraphs(string(rawBytes))
			return ChapterAudit{
				ChapterNum: chapterNum,
				RawParas:   len(rawParas),
				RawChars:   utf8.RuneCountInString(strings.Join(rawParas, "")),
				Status:     StatusPending,
				Detail:     "Not translated yet",
			}, ErrTransNotFound
		}
		return ChapterAudit{ChapterNum: chapterNum}, fmt.Errorf("read translated chapter %d: %w", chapterNum, err)
	}

	return ComputeAudit(chapterNum, string(rawBytes), string(htmlBytes)), nil
}

// ValidateRange audits chapters in range [start, end] and returns audit results.
func ValidateRange(rawDir, transDir string, start, end int) ([]ChapterAudit, error) {
	if start > end {
		return nil, fmt.Errorf("start chapter %d cannot exceed end chapter %d", start, end)
	}

	results := make([]ChapterAudit, 0, end-start+1)
	for ch := start; ch <= end; ch++ {
		rawPath := filepath.Join(rawDir, fmt.Sprintf("chapter_%d.txt", ch))
		htmlPath := filepath.Join(transDir, fmt.Sprintf("chapter_%d.html", ch))

		audit, err := ValidateChapter(rawPath, htmlPath, ch)
		if err != nil && !errors.Is(err, ErrRawNotFound) && !errors.Is(err, ErrTransNotFound) {
			return nil, err
		}
		results = append(results, audit)
	}

	return results, nil
}

// PrintAuditTable formats the audit results into a readable CLI table.
func PrintAuditTable(w io.Writer, results []ChapterAudit) {
	fmt.Fprintf(w, "\n%-8s | %-10s | %-12s | %-10s | %-10s | %-7s | %-8s | %s\n",
		"Chapter", "Raw Paras", "Trans Paras", "Raw Chars", "Eng Words", "W/C Rat", "Status", "Details")
	fmt.Fprintln(w, strings.Repeat("-", 95))

	for _, a := range results {
		statusStr := fmt.Sprintf("[%s]", a.Status)
		if a.Status == StatusPending {
			statusStr = "[TODO]"
		}

		fmt.Fprintf(w, "%-8d | %-10d | %-12d | %-10d | %-10d | %-7.3f | %-8s | %s\n",
			a.ChapterNum, a.RawParas, a.TransParas, a.RawChars, a.EnglishWords, a.WordCharRatio, statusStr, a.Detail)
	}
	fmt.Fprintln(w)
}
