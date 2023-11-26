package parser

import (
	"fmt"
	"os"
	"time"

	"github.com/Pensk/web-fetch/internal/fetcher"
	"golang.org/x/net/html"
)

type Metadata struct {
	NumLinks  int
	NumImages int
	LastFetch time.Time
}

func ParseMetadata(filename string) (Metadata, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Metadata{}, fmt.Errorf("error opening file %s: %w", filename, err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		return Metadata{}, fmt.Errorf("error parsing HTML file %s: %w", filename, err)
	}

	var metadata Metadata

	lastFetch, ok := fetcher.LastFetchedTime[filename]
	if !ok {
		lastFetch = time.Now()
	}
	metadata.LastFetch = lastFetch

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "a" {
				metadata.NumLinks++
			} else if n.Data == "img" {
				metadata.NumImages++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return metadata, nil
}
