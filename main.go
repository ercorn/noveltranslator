package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	ExitSuccess = 0
	ExitError   = 1
	ExitUsage   = 2
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	os.Exit(run(ctx, os.Args[1:], os.Stdout, os.Stderr))
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	logger := slog.New(slog.NewTextHandler(stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	if len(args) == 0 {
		printUsage(stdout)
		return ExitUsage
	}

	cmd := args[0]
	switch cmd {
	case "validate", "-validate", "--validate":
		return runValidate(args[1:], stdout, stderr, logger)
	case "scrape", "-scrape", "--scrape":
		return runScrape(ctx, args[1:], stderr, logger)
	case "chunk", "-chunk", "--chunk":
		return runChunk(args[1:], stdout, stderr, logger)
	case "help", "-h", "-help", "--help":
		printUsage(stdout)
		return ExitSuccess
	default:
		// Fallback for flat flag invocation like: go run . -validate -start=1173
		if len(args) > 0 && args[0][0] == '-' {
			return runLegacy(ctx, args, stdout, stderr, logger)
		}
		fmt.Fprintf(stderr, "unknown command: %q\n", cmd)
		printUsage(stderr)
		return ExitUsage
	}
}

func runValidate(args []string, stdout, stderr io.Writer, logger *slog.Logger) int {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	fs.SetOutput(stderr)

	start := fs.Int("start", 1164, "Start chapter number")
	end := fs.Int("end", 1220, "End chapter number")
	rawDir := fs.String("rawdir", "./raw", "Directory containing raw chapter files")
	transDir := fs.String("transdir", "./translated", "Directory containing translated HTML files")

	if err := fs.Parse(args); err != nil {
		return ExitUsage
	}

	results, err := ValidateRange(*rawDir, *transDir, *start, *end)
	if err != nil {
		logger.Error("validation failed", "error", err)
		return ExitError
	}

	PrintAuditTable(stdout, results)

	hasFail := false
	for _, r := range results {
		if r.Status == StatusFail {
			hasFail = true
			break
		}
	}
	if hasFail {
		return ExitError
	}
	return ExitSuccess
}

func runScrape(ctx context.Context, args []string, stderr io.Writer, logger *slog.Logger) int {
	fs := flag.NewFlagSet("scrape", flag.ContinueOnError)
	fs.SetOutput(stderr)

	start := fs.Int("start", 1164, "Start chapter number")
	end := fs.Int("end", 1220, "End chapter number")
	novelID := fs.String("novel", "727963", "Novel ID on source website")
	rawDir := fs.String("rawdir", "./raw", "Directory to store raw chapter files")

	if err := fs.Parse(args); err != nil {
		return ExitUsage
	}

	if err := ScrapeRange(ctx, *novelID, *start, *end, *rawDir, logger); err != nil {
		logger.ErrorContext(ctx, "scraping failed", "error", err)
		return ExitError
	}
	return ExitSuccess
}

func runChunk(args []string, stdout, stderr io.Writer, logger *slog.Logger) int {
	fs := flag.NewFlagSet("chunk", flag.ContinueOnError)
	fs.SetOutput(stderr)

	chapter := fs.Int("chapter", 1164, "Chapter number to chunk")
	chunkSize := fs.Int("size", 40, "Number of paragraphs per chunk")
	rawDir := fs.String("rawdir", "./raw", "Directory containing raw chapter files")

	if err := fs.Parse(args); err != nil {
		return ExitUsage
	}

	chunks, err := GetChapterChunks(*rawDir, *chapter, *chunkSize)
	if err != nil {
		logger.Error("chunking failed", "chapter", *chapter, "error", err)
		return ExitError
	}

	for _, ch := range chunks {
		fmt.Fprintf(stdout, "=== Chunk %d (Lines %d to %d of chapter %d) ===\n",
			ch.ChunkIndex, ch.StartLine, ch.EndLine, *chapter)
		for j, line := range ch.Lines {
			fmt.Fprintf(stdout, "[%d] %s\n", ch.StartLine+j, line)
		}
		fmt.Fprintln(stdout)
	}
	return ExitSuccess
}

func runLegacy(ctx context.Context, args []string, stdout, stderr io.Writer, logger *slog.Logger) int {
	fs := flag.NewFlagSet("noveltranslator", flag.ContinueOnError)
	fs.SetOutput(stderr)

	scrapeFlag := fs.Bool("scrape", false, "Scrape raw chapters")
	validateFlag := fs.Bool("validate", false, "Audit translated HTML files")
	chunkFlag := fs.Bool("chunk", false, "Inspect raw chapter chunks")
	chapterFlag := fs.Int("chapter", 1164, "Chapter number for chunking")
	chunkSizeFlag := fs.Int("size", 40, "Paragraph chunk size")
	startFlag := fs.Int("start", 1164, "Start chapter number")
	endFlag := fs.Int("end", 1220, "End chapter number")
	novelFlag := fs.String("novel", "727963", "Novel ID")
	rawDirFlag := fs.String("rawdir", "./raw", "Directory for raw files")
	transDirFlag := fs.String("transdir", "./translated", "Directory for translated files")

	if err := fs.Parse(args); err != nil {
		return ExitUsage
	}

	if *scrapeFlag {
		if err := ScrapeRange(ctx, *novelFlag, *startFlag, *endFlag, *rawDirFlag, logger); err != nil {
			logger.ErrorContext(ctx, "scraping failed", "error", err)
			return ExitError
		}
		return ExitSuccess
	}

	if *validateFlag {
		results, err := ValidateRange(*rawDirFlag, *transDirFlag, *startFlag, *endFlag)
		if err != nil {
			logger.Error("validation failed", "error", err)
			return ExitError
		}
		PrintAuditTable(stdout, results)
		return ExitSuccess
	}

	if *chunkFlag {
		chunks, err := GetChapterChunks(*rawDirFlag, *chapterFlag, *chunkSizeFlag)
		if err != nil {
			logger.Error("chunking failed", "error", err)
			return ExitError
		}
		for _, ch := range chunks {
			fmt.Fprintf(stdout, "=== Chunk %d (Lines %d to %d) ===\n", ch.ChunkIndex, ch.StartLine, ch.EndLine)
		}
		return ExitSuccess
	}

	printUsage(stdout)
	return ExitUsage
}

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "Novel Translator CLI")
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  noveltranslator validate [-start=N] [-end=N] [-rawdir=DIR] [-transdir=DIR]")
	fmt.Fprintln(w, "  noveltranslator scrape   [-start=N] [-end=N] [-novel=ID] [-rawdir=DIR]")
	fmt.Fprintln(w, "  noveltranslator chunk    [-chapter=N] [-size=N] [-rawdir=DIR]")
	fmt.Fprintln(w, "  noveltranslator help")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Legacy flag syntax (e.g. -validate -start=1173) is also supported for backward compatibility.")
}
