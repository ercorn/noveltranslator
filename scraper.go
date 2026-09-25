package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// ErrNoChaptersMatched indicates no chapters in the table of contents matched the requested range.
	ErrNoChaptersMatched = errors.New("no chapters matched requested range")
)

var (
	reLinkRegex      = regexp.MustCompile(`<a href="(/chapter/\w+/(\d+)\.html)"[^>]*>([^<]+)</a>`)
	reChapterNum     = regexp.MustCompile(`第(\d+)章`)
	reParagraph      = regexp.MustCompile(`(?s)<p>(.*?)</p>`)
	reNextPageLink   = regexp.MustCompile(`<a [^>]*href="(/chapter/[^"]+)"[^>]*>下一[頁页]</a>`)
	reHTMLTitle      = regexp.MustCompile(`<title>([^<]+)</title>`)
	reStripHTMLTags  = regexp.MustCompile(`<[^>]+>`)
)

type ChapterInfo struct {
	Number int
	Title  string
	URL    string
}

// sleepWithContext pauses execution for duration d or returns early if ctx is cancelled.
func sleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FetchTOC retrieves all chapter links and numbers from the novel table of contents page.
func FetchTOC(ctx context.Context, client *http.Client, novelID string) ([]ChapterInfo, error) {
	tocURL := fmt.Sprintf("https://www.uuread.tw/%s", novelID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tocURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create toc request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch toc: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch toc: unexpected status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read toc body: %w", err)
	}

	htmlContent := string(body)
	matches := reLinkRegex.FindAllStringSubmatch(htmlContent, -1)
	chapters := make([]ChapterInfo, 0, len(matches))
	seen := make(map[int]struct{}, len(matches))

	for _, m := range matches {
		href := m[1]
		title := strings.TrimSpace(m[3])

		numMatch := reChapterNum.FindStringSubmatch(title)
		if len(numMatch) < 2 {
			continue
		}
		num, err := strconv.Atoi(numMatch[1])
		if err != nil {
			continue
		}
		if _, exists := seen[num]; exists {
			continue
		}
		seen[num] = struct{}{}

		chapters = append(chapters, ChapterInfo{
			Number: num,
			Title:  title,
			URL:    "https://www.uuread.tw" + href,
		})
	}

	return chapters, nil
}

// FetchChapterRaw downloads all pages for a chapter (handling multi-page pagination) and returns cleaned paragraphs.
func FetchChapterRaw(ctx context.Context, client *http.Client, startURL string) (string, []string, error) {
	currentURL := startURL
	var allParagraphs []string
	var chapterTitle string

	for currentURL != "" {
		if err := ctx.Err(); err != nil {
			return "", nil, err
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, currentURL, nil)
		if err != nil {
			return "", nil, fmt.Errorf("create chapter request: %w", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

		resp, err := client.Do(req)
		if err != nil {
			return "", nil, fmt.Errorf("fetch chapter url %q: %w", currentURL, err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return "", nil, fmt.Errorf("read chapter response: %w", err)
		}

		htmlContent := string(body)

		if chapterTitle == "" {
			tMatch := reHTMLTitle.FindStringSubmatch(htmlContent)
			if len(tMatch) > 1 {
				parts := strings.Split(tMatch[1], "_")
				if len(parts) >= 2 {
					chapterTitle = parts[1]
				} else {
					chapterTitle = tMatch[1]
				}
			}
		}

		pMatches := reParagraph.FindAllStringSubmatch(htmlContent, -1)
		for _, pm := range pMatches {
			rawP := pm[1]
			cleanP := strings.TrimSpace(reStripHTMLTags.ReplaceAllString(rawP, ""))
			if cleanP == "" || strings.HasPrefix(cleanP, "uu") || strings.Contains(cleanP, "請記住本書首發域名") {
				continue
			}
			allParagraphs = append(allParagraphs, cleanP)
		}

		nextMatch := reNextPageLink.FindStringSubmatch(htmlContent)
		if len(nextMatch) > 1 {
			baseParsed, err := url.Parse(currentURL)
			if err != nil {
				return "", nil, fmt.Errorf("parse current url %q: %w", currentURL, err)
			}
			relParsed, err := url.Parse(nextMatch[1])
			if err != nil {
				return "", nil, fmt.Errorf("parse relative next url %q: %w", nextMatch[1], err)
			}
			resolvedURL := baseParsed.ResolveReference(relParsed).String()
			if resolvedURL == currentURL {
				currentURL = ""
			} else {
				currentURL = resolvedURL
				if err := sleepWithContext(ctx, 300*time.Millisecond); err != nil {
					return "", nil, err
				}
			}
		} else {
			currentURL = ""
		}
	}

	return chapterTitle, allParagraphs, nil
}

// ScrapeRange downloads chapters in range [start, end] and saves them to rawDir.
func ScrapeRange(ctx context.Context, novelID string, start, end int, rawDir string, logger *slog.Logger) error {
	if logger == nil {
		logger = slog.Default()
	}

	if err := os.MkdirAll(rawDir, 0o755); err != nil {
		return fmt.Errorf("create raw directory %q: %w", rawDir, err)
	}

	logger.InfoContext(ctx, "fetching novel table of contents", "novel_id", novelID)
	httpClient := &http.Client{Timeout: 30 * time.Second}

	allChapters, err := FetchTOC(ctx, httpClient, novelID)
	if err != nil {
		return fmt.Errorf("fetch toc: %w", err)
	}

	var targetChapters []ChapterInfo
	for _, ch := range allChapters {
		if ch.Number >= start && ch.Number <= end {
			targetChapters = append(targetChapters, ch)
		}
	}

	if len(targetChapters) == 0 {
		return fmt.Errorf("%w: range [%d, %d]", ErrNoChaptersMatched, start, end)
	}

	logger.InfoContext(ctx, "matched target chapters", "count", len(targetChapters), "start", start, "end", end)

	for i, ch := range targetChapters {
		if err := ctx.Err(); err != nil {
			return err
		}

		outPath := filepath.Join(rawDir, fmt.Sprintf("chapter_%d.txt", ch.Number))
		if _, err := os.Stat(outPath); err == nil {
			logger.InfoContext(ctx, "chapter already exists, skipping", "index", i+1, "total", len(targetChapters), "chapter", ch.Number)
			continue
		}

		logger.InfoContext(ctx, "downloading chapter", "index", i+1, "total", len(targetChapters), "chapter", ch.Number, "title", ch.Title)
		title, paras, err := FetchChapterRaw(ctx, httpClient, ch.URL)
		if err != nil {
			logger.ErrorContext(ctx, "failed downloading chapter", "chapter", ch.Number, "error", err)
			continue
		}

		fullContent := fmt.Sprintf("# %s\n\n%s\n", title, strings.Join(paras, "\n\n"))
		if err := os.WriteFile(outPath, []byte(fullContent), 0o644); err != nil {
			return fmt.Errorf("write chapter %d to %q: %w", ch.Number, outPath, err)
		}

		logger.InfoContext(ctx, "saved chapter", "chapter", ch.Number, "paragraphs", len(paras))
		if err := sleepWithContext(ctx, 400*time.Millisecond); err != nil {
			return err
		}
	}

	logger.InfoContext(ctx, "scraping completed successfully")
	return nil
}
