package fetcher

import (
	"io"
	"net/http"
)

type Fetcher struct{}

func (f *Fetcher) Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
