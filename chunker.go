package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RawChunk represents a bounded slice of raw chapter paragraphs.
type RawChunk struct {
	ChunkIndex int
	StartLine  int
	EndLine    int
	Lines      []string
}

// GetChapterChunks reads a raw chapter and splits non-blank lines into chunks of chunkSize.
func GetChapterChunks(rawDir string, chapterNum int, chunkSize int) ([]RawChunk, error) {
	rawPath := filepath.Join(rawDir, fmt.Sprintf("chapter_%d.txt", chapterNum))
	f, err := os.Open(rawPath)
	if err != nil {
		return nil, fmt.Errorf("open raw chapter %d: %w", chapterNum, err)
	}
	defer f.Close()

	var allLines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			allLines = append(allLines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan raw chapter %d: %w", chapterNum, err)
	}

	total := len(allLines)
	if total == 0 {
		return nil, fmt.Errorf("raw chapter %d is empty", chapterNum)
	}

	var chunks []RawChunk
	chunkIdx := 1
	for i := 0; i < total; i += chunkSize {
		end := i + chunkSize
		if end > total {
			end = total
		}
		chunks = append(chunks, RawChunk{
			ChunkIndex: chunkIdx,
			StartLine:  i + 1,
			EndLine:    end,
			Lines:      allLines[i:end],
		})
		chunkIdx++
	}

	return chunks, nil
}
