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
│   │   └── extractor.go  # Extracts links from parsed HTML
│   └── manager
│       └── manager.go    # Manages the crawling process
├── go.mod                # Module definition and dependencies
└── README.md             # Project documentation
```

## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd go-web-crawler
   ```

2. Initialize the Go module:
   ```
   go mod tidy
   ```

3. Build the application:
   ```
   go build -o crawler ./cmd/main.go
   ```

## Usage
To run the crawler, use the following command:
```
./crawler --url <start-url>
```

## Core Functionalities
- **Fetching**: The `Fetcher` struct sends HTTP GET requests to specified URLs and retrieves HTML content.
- **Parsing**: The `Parser` struct processes the HTML content and returns a structured representation of the document.
- **Extracting Links**: The `Extractor` struct extracts all links from the parsed HTML document.
- **Managing Crawls**: The `Manager` struct oversees the crawling process, tracking visited URLs and implementing crawling strategies.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.