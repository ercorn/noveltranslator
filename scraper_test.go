package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSleepWithContext(t *testing.T) {
	t.Run("completes when duration expires", func(t *testing.T) {
		ctx := context.Background()
		start := time.Now()
		err := sleepWithContext(ctx, 20*time.Millisecond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if elapsed := time.Since(start); elapsed < 15*time.Millisecond {
			t.Errorf("slept too short: %v", elapsed)
		}
	})

	t.Run("returns early when context is cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // immediately cancelled
		err := sleepWithContext(ctx, 1*time.Second)
		if err == nil {
			t.Fatal("expected context cancellation error, got nil")
		}
	})
}

func TestFetchChapterRaw(t *testing.T) {
	// Fake server serving a multi-page chapter
	var ts *httptest.Server
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/chapter/123/1.html":
			html := `<!DOCTYPE html>
<html>
<head><title>小說_第100章 測試標題</title></head>
<body>
  <p>第一段文本。</p>
  <p>第二段文本。</p>
  <a href="/chapter/123/2.html">下一頁</a>
</body>
</html>`
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(html))
		case "/chapter/123/2.html":
			html := `<!DOCTYPE html>
<html>
<head><title>小說_第100章 測試標題</title></head>
<body>
  <p>第三段文本。</p>
  <p>uu看書 www.uukanshu.com 廣告忽略</p>
</body>
</html>`
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(html))
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	title, paras, err := FetchChapterRaw(ctx, ts.Client(), ts.URL+"/chapter/123/1.html")
	if err != nil {
		t.Fatalf("FetchChapterRaw() error: %v", err)
	}

	if title != "第100章 測試標題" {
		t.Errorf("title = %q, want %q", title, "第100章 測試標題")
	}

	if len(paras) != 3 {
		t.Fatalf("got %d paragraphs, want 3 (ad should be filtered)", len(paras))
	}

	if paras[0] != "第一段文本。" || paras[1] != "第二段文本。" || paras[2] != "第三段文本。" {
		t.Errorf("unexpected paragraphs: %v", paras)
	}
}
