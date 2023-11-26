package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var LastFetchedTime map[string]time.Time

func FetchAndSave(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1] + ".html"

	var lastModified time.Time
	if fileInfo, err := os.Stat(filename); err == nil {
		lastModified = fileInfo.ModTime()
	} else {
		lastModified = time.Now()
	}

	LastFetchedTime[filename] = lastModified

	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("error creating file %s: %w", filename, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("error writing to file %s: %w", filename, err)
	}

	return filename, nil
}
