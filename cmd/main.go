package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Pensk/web-fetch/internal/fetcher"
	"github.com/Pensk/web-fetch/internal/parser"
)

func main() {
	metaFlag := flag.Bool("metadata", false, "Fetch and display metadata")
	flag.Parse()

	var urls []string

	if *metaFlag {
		urls = os.Args[2:]
	} else {
		urls = os.Args[1:]
	}

	fetcher.LastFetchedTime = make(map[string]time.Time)

	for _, url := range urls {
		filename, err := fetcher.FetchAndSave(url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if *metaFlag {
			metadata, err := parser.ParseMetadata(filename)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("site: %s\nnum_links: %d\nimages: %d\nlast_fetch: %s\n", url, metadata.NumLinks, metadata.NumImages, metadata.LastFetch.Format(time.RFC1123))
		}
	}
}
