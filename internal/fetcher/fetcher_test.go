package fetcher

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestFetchAndSave(t *testing.T) {
	html := "<html><body>OK</body></html>"

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(html))
	}))
	serverURL := server.URL + "/OK"

	defer server.Close()

	LastFetchedTime = make(map[string]time.Time)

	filename, err := FetchAndSave(serverURL)
	if err != nil {
		t.Errorf("FetchAndSave failed: %v", err)
	}

	// Check the file was created
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("File was not created")
	}

	// Check the contents of the file
	contents, err := os.ReadFile("OK.html")
	if err != nil {
		t.Errorf("Failed to read file: %v", err)
	}
	if string(contents) != html {
		t.Errorf("File contents did not match server output")
	}

	os.Remove("OK.html")
}
