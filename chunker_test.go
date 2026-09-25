package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetChapterChunks(t *testing.T) {
	tempDir := t.TempDir()

	// Create valid sample chapter
	sampleLines := []string{
		"# 第1章 測試",
		"第1段",
		"第2段",
		"第3段",
		"第4段",
		"第5段",
	}
	ch1Path := filepath.Join(tempDir, "chapter_1.txt")
	if err := os.WriteFile(ch1Path, []byte(strings.Join(sampleLines, "\n\n")), 0o644); err != nil {
		t.Fatalf("failed to write test fixture: %v", err)
	}

	// Create empty sample chapter
	ch2Path := filepath.Join(tempDir, "chapter_2.txt")
	if err := os.WriteFile(ch2Path, []byte("# 標題只有一行\n\n"), 0o644); err != nil {
		t.Fatalf("failed to write empty test fixture: %v", err)
	}

	tests := []struct {
		name       string
		chapterNum int
		chunkSize  int
		wantChunks int
		wantErr    error
	}{
		{
			name:       "valid chunks of size 2",
			chapterNum: 1,
			chunkSize:  2,
			wantChunks: 3, // 5 paras / 2 = 3 chunks (2, 2, 1)
			wantErr:    nil,
		},
		{
			name:       "valid single chunk when size exceeds total",
			chapterNum: 1,
			chunkSize:  10,
			wantChunks: 1,
			wantErr:    nil,
		},
		{
			name:       "invalid chunk size zero",
			chapterNum: 1,
			chunkSize:  0,
			wantChunks: 0,
			wantErr:    ErrInvalidChunkSize,
		},
		{
			name:       "invalid chunk size negative",
			chapterNum: 1,
			chunkSize:  -5,
			wantChunks: 0,
			wantErr:    ErrInvalidChunkSize,
		},
		{
			name:       "empty chapter error",
			chapterNum: 2,
			chunkSize:  5,
			wantChunks: 0,
			wantErr:    ErrEmptyChapter,
		},
		{
			name:       "missing chapter file error",
			chapterNum: 999,
			chunkSize:  5,
			wantChunks: 0,
			wantErr:    ErrRawNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chunks, err := GetChapterChunks(tempDir, tt.chapterNum, tt.chunkSize)
			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("got error %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(chunks) != tt.wantChunks {
				t.Errorf("got %d chunks, want %d", len(chunks), tt.wantChunks)
			}

			// Verify contiguous line indices
			var totalLines int
			for i, ch := range chunks {
				if ch.ChunkIndex != i+1 {
					t.Errorf("chunk[%d].ChunkIndex = %d, want %d", i, ch.ChunkIndex, i+1)
				}
				totalLines += len(ch.Lines)
			}
			if totalLines != 5 {
				t.Errorf("total lines in chunks = %d, want 5", totalLines)
			}
		})
	}
}
