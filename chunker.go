package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	// ErrInvalidChunkSize indicates the requested chunk size is non-positive.
	ErrInvalidChunkSize = errors.New("chunk size must be greater than zero")
	// ErrEmptyChapter indicates the raw chapter has no paragraph content.
	ErrEmptyChapter = errors.New("raw chapter has no valid paragraphs")
)

// RawChunk represents a bounded slice of raw chapter paragraphs.
type RawChunk struct {
	ChunkIndex int
	StartLine  int
	EndLine    int
	Lines      []string
}

// GetChapterChunks reads a raw chapter file and partitions its paragraphs into chunks of chunkSize.
func GetChapterChunks(rawDir string, chapterNum, chunkSize int) ([]RawChunk, error) {
	if chunkSize <= 0 {
		return nil, ErrInvalidChunkSize
	}

	rawPath := filepath.Join(rawDir, fmt.Sprintf("chapter_%d.txt", chapterNum))
	rawBytes, err := os.ReadFile(rawPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("read chapter %d: %w", chapterNum, ErrRawNotFound)
		}
		return nil, fmt.Errorf("read chapter %d: %w", chapterNum, err)
	}

	paragraphs := ExtractRawParagraphs(string(rawBytes))
	total := len(paragraphs)
	if total == 0 {
		return nil, fmt.Errorf("chapter %d: %w", chapterNum, ErrEmptyChapter)
	}

	numChunks := (total + chunkSize - 1) / chunkSize
	chunks := make([]RawChunk, 0, numChunks)

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
			Lines:      paragraphs[i:end],
		})
		chunkIdx++
	}

	return chunks, nil
}
