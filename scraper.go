package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ChapterInfo struct {
	Number int
	Title  string
	URL    string
}

// FetchTOC retrieves all chapter links and numbers from the novel TOC page.
func FetchTOC(novelID string) ([]ChapterInfo, error) {
	tocURL := fmt.Sprintf("https://www.uuread.tw/%s", novelID)
	req, err := http.NewRequest("GET", tocURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	htmlContent := string(body)
	reLink := regexp.MustCompile(`<a href="(/chapter/` + novelID + `/(\d+)\.html)"[^>]*>([^<]+)</a>`)
	reNum := regexp.MustCompile(`第(\d+)章`)

	matches := reLink.FindAllStringSubmatch(htmlContent, -1)
	var chapters []ChapterInfo
	seen := make(map[int]bool)

	for _, m := range matches {
		href := m[1]
		title := strings.TrimSpace(m[3])

		numMatch := reNum.FindStringSubmatch(title)
		if len(numMatch) < 2 {
			continue
		}
		num, err := strconv.Atoi(numMatch[1])
		if err != nil || seen[num] {
			continue
		}
		seen[num] = true

		fullURL := "https://www.uuread.tw" + href
		chapters = append(chapters, ChapterInfo{
			Number: num,
			Title:  title,
			URL:    fullURL,
		})
	}

	return chapters, nil
}

// FetchChapterRaw downloads all pages for a chapter (including _2.html etc.) and returns the full text.
func FetchChapterRaw(client *http.Client, startURL string) (string, []string, error) {
	currentURL := startURL
	var allParagraphs []string
	var chapterTitle string

	reParagraph := regexp.MustCompile(`(?s)<p>(.*?)</p>`)
	reNextPage := regexp.MustCompile(`<a [^>]*href="(/chapter/[^"]+)"[^>]*>下一[頁页]</a>`)
	reTitle := regexp.MustCompile(`<title>([^<]+)</title>`)
	reTag := regexp.MustCompile(`<[^>]+>`)

	for currentURL != "" {
		req, err := http.NewRequest("GET", currentURL, nil)
		if err != nil {
			return "", nil, err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

		resp, err := client.Do(req)
		if err != nil {
			return "", nil, err
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return "", nil, err
		}

		htmlContent := string(body)

		if chapterTitle == "" {
			tMatch := reTitle.FindStringSubmatch(htmlContent)
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
			cleanP := strings.TrimSpace(reTag.ReplaceAllString(rawP, ""))
			if cleanP == "" || strings.HasPrefix(cleanP, "uu") || strings.Contains(cleanP, "請記住本書首發域名") {
				continue
			}
			allParagraphs = append(allParagraphs, cleanP)
		}

		// Check for multi-part pagination
		nextMatch := reNextPage.FindStringSubmatch(htmlContent)
		if len(nextMatch) > 1 {
			nextURL := "https://www.uuread.tw" + nextMatch[1]
			// Avoid infinite loop if next points to itself
			if nextURL == currentURL {
				currentURL = ""
			} else {
				currentURL = nextURL
				time.Sleep(300 * time.Millisecond) // Polite delay
			}
		} else {
			currentURL = ""
		}
	}

	return chapterTitle, allParagraphs, nil
}

// ScrapeRange downloads chapters in range [start, end] and saves them to rawDir.
func ScrapeRange(novelID string, start, end int, rawDir string) error {
	if err := os.MkdirAll(rawDir, 0755); err != nil {
		return err
	}

	fmt.Printf("Fetching Table of Contents for novel %s...\n", novelID)
	allChapters, err := FetchTOC(novelID)
	if err != nil {
		return fmt.Errorf("failed to fetch TOC: %w", err)
	}

	// Filter requested range
	var targetChapters []ChapterInfo
	for _, ch := range allChapters {
		if ch.Number >= start && ch.Number <= end {
			targetChapters = append(targetChapters, ch)
		}
	}

	fmt.Printf("Found %d chapters in range %d - %d.\n", len(targetChapters), start, end)

	httpClient := &http.Client{Timeout: 30 * time.Second}

	for i, ch := range targetChapters {
		outPath := filepath.Join(rawDir, fmt.Sprintf("chapter_%d.txt", ch.Number))

		// Check if already downloaded
		if _, err := os.Stat(outPath); err == nil {
			fmt.Printf("[%d/%d] Chapter %d already exists, skipping.\n", i+1, len(targetChapters), ch.Number)
			continue
		}

		fmt.Printf("[%d/%d] Downloading Chapter %d: %s...\n", i+1, len(targetChapters), ch.Number, ch.Title)
		title, paras, err := FetchChapterRaw(httpClient, ch.URL)
		if err != nil {
			fmt.Printf("Error downloading chapter %d: %v\n", ch.Number, err)
			continue
		}

		fullContent := fmt.Sprintf("# %s\n\n%s\n", title, strings.Join(paras, "\n\n"))
		if err := os.WriteFile(outPath, []byte(fullContent), 0644); err != nil {
			return fmt.Errorf("failed to write raw chapter %d: %w", ch.Number, err)
		}

		fmt.Printf("Saved Chapter %d (%d paragraphs)\n", ch.Number, len(paras))
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Scraping completed!")
	return nil
}
