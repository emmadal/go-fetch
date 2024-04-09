package pkg

import (
	"fmt"
	"net/http"
	"strings"
)

func getFileName(url string) string {
	// Get the last part of a URL.
	splits := strings.Split(url, "/")
	return splits[len(splits)-1]
}

func info(response *http.Response, filename string) {
	fmt.Printf("Date: %s\n", response.Header.Get("Date"))
	fmt.Printf("Location... %s\n", response.Request.URL)
	fmt.Printf("HTTP request sent, awaiting response... %s\n", response.Status)
	fmt.Printf("Length: %d [application/octet-stream]\n", response.ContentLength)
	fmt.Printf("Saving to: `%s`\n", filename)
}