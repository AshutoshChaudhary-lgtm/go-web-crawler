# Go Web Crawler

## Overview
This project implements a web crawler in Go that fetches web pages, parses HTML content, extracts links, and manages the crawling process. It is designed to be efficient and easy to use.

## Project Structure
```
go-web-crawler
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── fetcher
│   │   └── fetcher.go   # Fetches HTML content from URLs
│   ├── parser
│   │   └── parser.go    # Parses HTML content
│   ├── extractor
│   │   └── extractor.go # Extracts links from parsed HTML
│   └── manager
│       └── manager.go   # Manages the crawling process
├── go.mod               # Module definition and dependencies
├── go.sum               # Checksums for module dependencies
└── README.md            # Project documentation
```

## Setup Instructions
1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd go-web-crawler
   ```

2. Initialize the Go module:
   ```sh
   go mod tidy
   ```

3. Build the application:
   ```sh
   go build -o crawler ./cmd/main.go
   ```

## Usage
To run the crawler, use the following command:
```sh
./crawler --url <start-url> --depth <depth> --concurrency <concurrency> --user-agent <user-agent> --timeout <timeout> --output <output-file> [--impolite]
```

### Example
```sh
./crawler --url https://example.com --depth 2 --concurrency 5 --user-agent "MyCrawler" --timeout 15s --output results.txt
```

### Example (Ignoring robots.txt)
```sh
./crawler --url https://example.com --depth 2 --concurrency 5 --user-agent "MyCrawler" --timeout 15s --output results.txt --impolite
```

## Core Functionalities
- **Fetching**: The `Fetcher` struct sends HTTP GET requests to specified URLs and retrieves HTML content.
- **Parsing**: The `Parser` struct processes the HTML content and returns a structured representation of the document.
- **Extracting Links**: The `Extractor` struct extracts all links from the parsed HTML document.
- **Managing Crawls**: The `Manager` struct oversees the crawling process, tracking visited URLs and implementing crawling strategies.
- **Respecting robots.txt**: The `Manager` struct checks the `robots.txt` file of the website to ensure compliance with the site's crawling policies, unless the `--impolite` flag is used.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.