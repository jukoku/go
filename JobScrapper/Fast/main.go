package main

import "github.com/go/JobScrapper/Fast/scrapper"

func main() {
	// Job <- input keyword what you want ex) Python, Java, etc...
	scrapper.Scrape("Golang")
}
