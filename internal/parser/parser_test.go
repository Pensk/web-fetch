package parser

import (
	"os"
	"testing"
	"time"

	"github.com/Pensk/web-fetch/internal/fetcher"
)

func TestParseMetadata(t *testing.T) {
	html := `<html><body>
        <a href="https://example.com">Link</a>
        <img src="image.jpg">
    </body></html>`

	file, err := os.CreateTemp("", "*.html")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(html))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	fetcher.LastFetchedTime = make(map[string]time.Time)
	fetcher.LastFetchedTime[file.Name()] = time.Now()

	metadata, err := ParseMetadata(file.Name())
	if err != nil {
		t.Errorf("ParseMetadata failed: %v", err)
	}

	if metadata.NumLinks != 1 {
		t.Errorf("NumLinks was incorrect, got: %d, want: %d", metadata.NumLinks, 1)
	}
	if metadata.NumImages != 1 {
		t.Errorf("NumImages was incorrect, got: %d, want: %d", metadata.NumImages, 1)
	}
	if metadata.LastFetch != fetcher.LastFetchedTime[file.Name()] {
		t.Errorf("LastFetch was incorrect, got: %v, want: %v", metadata.LastFetch, fetcher.LastFetchedTime[file.Name()])
	}
}
