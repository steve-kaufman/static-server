package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var staticDirectory string

func main() {
	staticDirectory = os.Getenv("STATIC_DIR")
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fileLocation, err := getFile(staticDirectory, r.URL.Path)

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404 not found")
	}

	http.ServeFile(w, r, fileLocation)
}

func getFile(rootDir string, requestedFile string) (string, error) {
	fileLocation := filepath.Join(rootDir, requestedFile)

	if fileExists(fileLocation) {
		return fileLocation, nil
	}

	if fileExists(fmt.Sprintf("%s.html", fileLocation)) {
		return fmt.Sprintf("%s.html", fileLocation), nil
	}

	return "", errors.New("no such file")
}

func fileExists(fileLocation string) bool {
	if _, err := os.Stat(fileLocation); err != nil {
		return false
	}
	return true
}
