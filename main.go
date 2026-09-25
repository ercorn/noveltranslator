package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	scrapeFlag := flag.Bool("scrape", false, "Scrape raw chapters from uuread.tw")
	startFlag := flag.Int("start", 1165, "Start chapter number")
	endFlag := flag.Int("end", 1220, "End chapter number")
	novelFlag := flag.String("novel", "727963", "Novel ID on uuread.tw")
	rawDirFlag := flag.String("rawdir", "./raw", "Directory to store raw chapter files")

	flag.Parse()

	if *scrapeFlag {
		fmt.Printf("=== Starting Scraper for Novel %s (Chapters %d to %d) ===\n", *novelFlag, *startFlag, *endFlag)
		if err := ScrapeRange(*novelFlag, *startFlag, *endFlag, *rawDirFlag); err != nil {
			log.Fatalf("Scraping failed: %v\n", err)
		}
		return
	}

	fmt.Println("Novel Translator CLI")
	fmt.Println("Usage:")
	fmt.Println("  go run . -scrape -start=1165 -end=1220")
	flag.PrintDefaults()
}
