package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func FetchAndSave(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1] + ".html"

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filename, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", filename, err)
	}

	return nil
}
