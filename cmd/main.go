package main

import (
	"fmt"
	"os"

	"github.com/Pensk/web-fetch/internal/fetcher"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		if err := fetcher.FetchAndSave(url); err != nil {
			fmt.Println(err)
		}
	}
}
