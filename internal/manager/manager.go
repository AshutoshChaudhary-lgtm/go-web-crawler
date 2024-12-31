package manager

import (
	"fmt"
	"go-web-crawler/internal/extractor"
	"go-web-crawler/internal/parser"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type Manager struct {
	visited     map[string]bool
	mu          sync.Mutex
	parser      *parser.Parser
	extractor   *extractor.Extractor
	depth       int
	concurrency int
	userAgent   string
	timeout     time.Duration
	outputFile  string
	outputMu    sync.Mutex
}

func NewManager(depth, concurrency int, userAgent string, timeout time.Duration, outputFile string) *Manager {
	return &Manager{
		visited:     make(map[string]bool),
		parser:      &parser.Parser{},
		extractor:   &extractor.Extractor{},
		depth:       depth,
		concurrency: concurrency,
		userAgent:   userAgent,
		timeout:     timeout,
		outputFile:  outputFile,
	}
}

func (m *Manager) ManageCrawl(startURL string, currentDepth int) error {
	if currentDepth > m.depth {
		return nil
	}

	m.mu.Lock()
	if m.visited[startURL] {
		m.mu.Unlock()
		return nil
	}
	m.visited[startURL] = true
	m.mu.Unlock()

	fmt.Printf("Visiting: %s\n", startURL)

	content, err := m.fetchContent(startURL)
	if err != nil {
		return fmt.Errorf("error fetching: %w", err)
	}

	doc, err := m.parser.Parse(content)
	if err != nil {
		return fmt.Errorf("error parsing: %w", err)
	}

	links, err := m.extractor.ExtractLinks(doc)
	if err != nil {
		return fmt.Errorf("error extracting links: %w", err)
	}

	baseURL, err := url.Parse(startURL)
	if err != nil {
		return fmt.Errorf("error parsing base URL: %w", err)
	}

	for _, link := range links {
		absoluteURL := resolveURL(link, baseURL)
		if err := m.ManageCrawl(absoluteURL, currentDepth+1); err != nil {
			fmt.Println("Error during crawling:", err)
		}
	}

	m.saveURL(startURL)

	return nil
}

func (m *Manager) fetchContent(url string) (string, error) {
	client := &http.Client{
		Timeout: m.timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", m.userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func resolveURL(link string, baseURL *url.URL) string {
	parsedLink, err := url.Parse(link)
	if err != nil {
		return link
	}
	return baseURL.ResolveReference(parsedLink).String()
}

func (m *Manager) saveURL(url string) {
	m.outputMu.Lock()
	defer m.outputMu.Unlock()

	file, err := os.OpenFile(m.outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening output file: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(url + "\n"); err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
	}
}
