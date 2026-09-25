package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

// ChapterAudit holds verification metrics for a single chapter.
type ChapterAudit struct {
	ChapterNum     int
	RawParas       int
	TransParas     int
	RawChars       int
	EnglishWords   int
	WordCharRatio  float64
	ParaRatio      float64
	EmptyParagraph int
	Status         string
	Detail         string
}

var (
	pTagRegex     = regexp.MustCompile(`(?s)<p>(.*?)</p>`)
	tagStripRegex = regexp.MustCompile(`<[^>]+>`)
	articleRegex  = regexp.MustCompile(`(?s)<article[^>]*id="chapter-content"[^>]*>(.*?)</article>`)
)

// ValidateRange audits translated HTML files against raw source text for the specified chapter range.
func ValidateRange(rawDir, transDir string, start, end int) ([]ChapterAudit, error) {
	var results []ChapterAudit

	for ch := start; ch <= end; ch++ {
		rawPath := filepath.Join(rawDir, fmt.Sprintf("chapter_%d.txt", ch))
		htmlPath := filepath.Join(transDir, fmt.Sprintf("chapter_%d.html", ch))

		audit := ChapterAudit{ChapterNum: ch}

		rawBytes, err := os.ReadFile(rawPath)
		if err != nil {
			audit.Status = "NO_RAW"
			audit.Detail = "Raw source file missing"
			results = append(results, audit)
			continue
		}

		rawLines := strings.Split(string(rawBytes), "\n")
		var nonBlankRaw []string
		for _, line := range rawLines {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
				nonBlankRaw = append(nonBlankRaw, trimmed)
			}
		}

		audit.RawParas = len(nonBlankRaw)
		audit.RawChars = utf8.RuneCountInString(strings.Join(nonBlankRaw, ""))

		htmlBytes, err := os.ReadFile(htmlPath)
		if err != nil {
			audit.Status = "PENDING"
			audit.Detail = "Not translated yet"
			results = append(results, audit)
			continue
		}

		htmlStr := string(htmlBytes)
		articleMatch := articleRegex.FindStringSubmatch(htmlStr)
		targetHTML := htmlStr
		if len(articleMatch) > 1 {
			targetHTML = articleMatch[1]
		}

		matches := pTagRegex.FindAllStringSubmatch(targetHTML, -1)
		audit.TransParas = len(matches)

		var allWords []string
		for _, m := range matches {
			content := strings.TrimSpace(m[1])
			if content == "" {
				audit.EmptyParagraph++
				continue
			}
			plain := tagStripRegex.ReplaceAllString(content, " ")
			words := strings.Fields(plain)
			allWords = append(allWords, words...)
		}
		audit.EnglishWords = len(allWords)

		if audit.RawChars > 0 {
			audit.WordCharRatio = float64(audit.EnglishWords) / float64(audit.RawChars)
		}
		if audit.RawParas > 0 {
			audit.ParaRatio = float64(audit.TransParas) / float64(audit.RawParas)
		}

		if audit.EmptyParagraph > 0 {
			audit.Status = "FAIL"
			audit.Detail = fmt.Sprintf("Contains %d empty <p> tags", audit.EmptyParagraph)
		} else if audit.WordCharRatio < 0.50 || audit.ParaRatio < 0.60 {
			audit.Status = "FAIL"
			audit.Detail = fmt.Sprintf("Severely condensed (ratio: %.2f, paras: %d/%d)", audit.WordCharRatio, audit.TransParas, audit.RawParas)
		} else if audit.WordCharRatio < 0.58 || audit.ParaRatio < 0.75 {
			audit.Status = "WARN"
			audit.Detail = fmt.Sprintf("Partially condensed (ratio: %.2f)", audit.WordCharRatio)
		} else {
			audit.Status = "PASS"
			audit.Detail = "Full fidelity verified"
		}

		results = append(results, audit)
	}

	return results, nil
}

// PrintAuditTable formats the audit results into a readable CLI table.
func PrintAuditTable(results []ChapterAudit) {
	fmt.Printf("\n%-8s | %-10s | %-12s | %-10s | %-10s | %-7s | %-8s | %s\n",
		"Chapter", "Raw Paras", "Trans Paras", "Raw Chars", "Eng Words", "W/C Rat", "Status", "Details")
	fmt.Println(strings.Repeat("-", 95))

	for _, a := range results {
		statusStr := a.Status
		switch a.Status {
		case "PASS":
			statusStr = "[PASS]"
		case "FAIL":
			statusStr = "[FAIL]"
		case "WARN":
			statusStr = "[WARN]"
		case "PENDING":
			statusStr = "[TODO]"
		}

		fmt.Printf("%-8d | %-10d | %-12d | %-10d | %-10d | %-7.3f | %-8s | %s\n",
			a.ChapterNum, a.RawParas, a.TransParas, a.RawChars, a.EnglishWords, a.WordCharRatio, statusStr, a.Detail)
	}
	fmt.Println()
}
