package pkg

import (
	"io"
	"log"
	"net/http"
	"os"
	"github.com/schollz/progressbar/v3"
)

// Fetcher download the file in the given URL
func Fetcher(url string) {
	response, err := http.Get(url) // Download file from URL and ignore response
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	filename := getFileName(url)
	info(response, filename)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	bar := progressbar.DefaultBytes(response.ContentLength, "Downloading")
	io.Copy(io.MultiWriter(file, bar), response.Body)
}
