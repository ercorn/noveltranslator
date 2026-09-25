package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestExtractRawParagraphs(t *testing.T) {
	tests := []struct {
		name      string
		rawText   string
		wantParas []string
	}{
		{
			name: "skips titles and blank lines",
			rawText: "# 第100章 測試\n\n第一段文字。\n\n第二段文字。\n\n# 另一個標題\n第三段文字。\n",
			wantParas: []string{
				"第一段文字。",
				"第二段文字。",
				"第三段文字。",
			},
		},
		{
			name:      "empty input",
			rawText:   "   \n\n\n",
			wantParas: []string{},
		},
		{
			name:      "trims whitespace",
			rawText:   "   帶有前後空格的段落。   \n\t帶有製表符的段落。\t",
			wantParas: []string{
				"帶有前後空格的段落。",
				"帶有製表符的段落。",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractRawParagraphs(tt.rawText)
			if len(got) != len(tt.wantParas) {
				t.Fatalf("ExtractRawParagraphs() returned %d paras, want %d", len(got), len(tt.wantParas))
			}
			for i := range got {
				if got[i] != tt.wantParas[i] {
					t.Errorf("para[%d] = %q, want %q", i, got[i], tt.wantParas[i])
				}
			}
		})
	}
}

func TestExtractTranslatedParagraphs(t *testing.T) {
	tests := []struct {
		name           string
		htmlContent    string
		wantParas      []string
		wantEmptyCount int
	}{
		{
			name: "extracts from article container",
			htmlContent: "<!DOCTYPE html>\n<html>\n<body>\n  <article id=\"chapter-content\">\n    <p>First translated paragraph.</p>\n    <p>Second translated paragraph with <em>emphasis</em>.</p>\n  </article>\n</body>\n</html>",
			wantParas: []string{
				"First translated paragraph.",
				"Second translated paragraph with  emphasis .",
			},
			wantEmptyCount: 0,
		},
		{
			name: "detects empty paragraphs",
			htmlContent: "<article id=\"chapter-content\">\n  <p>Valid paragraph.</p>\n  <p></p>\n  <p>   </p>\n  <p>Another valid paragraph.</p>\n</article>",
			wantParas: []string{
				"Valid paragraph.",
				"Another valid paragraph.",
			},
			wantEmptyCount: 2,
		},
		{
			name: "falls back to whole document if article tag absent",
			htmlContent: "<div>\n  <p>Fallback paragraph one.</p>\n  <p>Fallback paragraph two.</p>\n</div>",
			wantParas: []string{
				"Fallback paragraph one.",
				"Fallback paragraph two.",
			},
			wantEmptyCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParas, gotEmpty := ExtractTranslatedParagraphs(tt.htmlContent)
			if gotEmpty != tt.wantEmptyCount {
				t.Errorf("got emptyCount = %d, want %d", gotEmpty, tt.wantEmptyCount)
			}
			if len(gotParas) != len(tt.wantParas) {
				t.Fatalf("got %d paras, want %d", len(gotParas), len(tt.wantParas))
			}
			for i := range gotParas {
				if strings.TrimSpace(gotParas[i]) != strings.TrimSpace(tt.wantParas[i]) {
					t.Errorf("para[%d] = %q, want %q", i, gotParas[i], tt.wantParas[i])
				}
			}
		})
	}
}

func TestComputeAudit(t *testing.T) {
	tests := []struct {
		name        string
		rawText     string
		htmlContent string
		wantStatus  AuditStatus
	}{
		{
			name: "full fidelity pass",
			rawText: "第一段中文文本。\n第二段中文文本。",
			htmlContent: "<article id=\"chapter-content\">\n  <p>This is the first English translated paragraph with enough words to meet the ratio threshold.</p>\n  <p>This is the second English translated paragraph that also has sufficient word count.</p>\n</article>",
			wantStatus: StatusPass,
		},
		{
			name: "empty paragraph causes fail",
			rawText: "第一段中文。\n第二段中文。",
			htmlContent: "<article id=\"chapter-content\">\n  <p>First paragraph.</p>\n  <p></p>\n</article>",
			wantStatus: StatusFail,
		},
		{
			name: "severely low word ratio causes fail",
			rawText: strings.Repeat("這是一段非常非常長的中文字符內容。", 20),
			htmlContent: "<article id=\"chapter-content\">\n  <p>Short.</p>\n</article>",
			wantStatus: StatusFail,
		},
		{
			name: "borderline ratio causes warn",
			rawText: strings.Repeat("一二三四五六七八九十", 10), // 100 chars
			htmlContent: "<article id=\"chapter-content\"><p>" + strings.Repeat("word ", 55) + "</p></article>", // 55 words -> 0.55 ratio
			wantStatus: StatusWarn,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			audit := ComputeAudit(100, tt.rawText, tt.htmlContent)
			if audit.Status != tt.wantStatus {
				t.Errorf("ComputeAudit() status = %v, want %v (detail: %s)", audit.Status, tt.wantStatus, audit.Detail)
			}
		})
	}
}

func TestPrintAuditTable(t *testing.T) {
	results := []ChapterAudit{
		{
			ChapterNum:    1173,
			RawParas:      193,
			TransParas:    193,
			RawChars:      7529,
			EnglishWords:  5471,
			WordCharRatio: 0.727,
			Status:        StatusPass,
			Detail:        "Full fidelity verified",
		},
		{
			ChapterNum:    1174,
			RawParas:      199,
			TransParas:    199,
			RawChars:      7600,
			EnglishWords:  6091,
			WordCharRatio: 0.801,
			Status:        StatusPass,
			Detail:        "Full fidelity verified",
		},
	}

	var buf bytes.Buffer
	PrintAuditTable(&buf, results)

	output := buf.String()
	if !strings.Contains(output, "1173") || !strings.Contains(output, "1174") {
		t.Errorf("expected table to contain chapter numbers, got:\n%s", output)
	}
	if !strings.Contains(output, "[PASS]") {
		t.Errorf("expected table to contain [PASS], got:\n%s", output)
	}
}
