package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	scrapeFlag := flag.Bool("scrape", false, "Scrape raw chapters from uuread.tw")
	validateFlag := flag.Bool("validate", false, "Audit translated HTML files against raw text")
	startFlag := flag.Int("start", 1164, "Start chapter number")
	endFlag := flag.Int("end", 1220, "End chapter number")
	novelFlag := flag.String("novel", "727963", "Novel ID on uuread.tw")
	rawDirFlag := flag.String("rawdir", "./raw", "Directory to store raw chapter files")
	transDirFlag := flag.String("transdir", "./translated", "Directory to store translated HTML files")

	flag.Parse()

	if *scrapeFlag {
		fmt.Printf("=== Starting Scraper for Novel %s (Chapters %d to %d) ===\n", *novelFlag, *startFlag, *endFlag)
		if err := ScrapeRange(*novelFlag, *startFlag, *endFlag, *rawDirFlag); err != nil {
			log.Fatalf("Scraping failed: %v\n", err)
		}
		return
	}

	if *validateFlag {
		fmt.Printf("=== Validating Chapters %d to %d ===\n", *startFlag, *endFlag)
		results, err := ValidateRange(*rawDirFlag, *transDirFlag, *startFlag, *endFlag)
		if err != nil {
			log.Fatalf("Validation error: %v\n", err)
		}
		PrintAuditTable(results)
		return
	}

	fmt.Println("Novel Translator CLI")
	fmt.Println("Usage:")
	fmt.Println("  go run . -scrape -start=1164 -end=1220")
	fmt.Println("  go run . -validate -start=1164 -end=1170")
	flag.PrintDefaults()
}
