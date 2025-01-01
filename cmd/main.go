package main

import (
	"flag"
	"fmt"
	"go-web-crawler/internal/manager"
	"log"
	"os"
	"time"
)

func main() {
	startURL := flag.String("url", "", "The start URL for the crawler")
	depth := flag.Int("depth", 1, "The maximum depth to crawl")
	concurrency := flag.Int("concurrency", 1, "The number of concurrent requests")
	userAgent := flag.String("user-agent", "Go-Web-Crawler", "The User-Agent string for HTTP requests")
	timeout := flag.Duration("timeout", 10*time.Second, "The timeout for HTTP requests")
	outputFile := flag.String("output", "output.txt", "The file to save crawled URLs")
	impolite := flag.Bool("impolite", false, "Ignore robots.txt rules")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *startURL == "" {
		flag.Usage()
		os.Exit(1)
	}

	crawler := manager.NewManager(*depth, *concurrency, *userAgent, *timeout, *outputFile)

	if !*impolite && !crawler.CheckRobotsTxt(*startURL) {
		log.Fatalf("Crawling disallowed by robots.txt")
	}

	err := crawler.ManageCrawl(*startURL, 0)
	if err != nil {
		log.Fatalf("Error during crawling: %v", err)
	}

	fmt.Println("Crawling completed successfully.")
}
