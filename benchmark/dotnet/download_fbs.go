package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// Run as CGO_ENABLED=0 go run download_fbs.go

type File struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

func main() {
	resp, err := http.Get(`https://api.github.com/repos/google/flatbuffers/contents/net/FlatBuffers?ref=master`)
	if err != nil {
		panic(err)
	}

	var data []File
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	var wait sync.WaitGroup
	for _, entry := range data {
		if entry.DownloadURL == "" {
			continue
		}
		wait.Add(1)
		go func(f File) {
			defer wait.Done()
			resp, err := http.Get(f.DownloadURL)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			os.MkdirAll(filepath.Dir(f.Path), 0755)
			out, err := os.Create(f.Path)
			if err != nil {
				panic(err)
			}
			defer out.Close()
			if _, err := io.Copy(out, resp.Body); err != nil {
				panic(err)
			}
		}(entry)
	}

	wait.Wait()
}
