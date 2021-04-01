package main

import (
	"net/http"
	"os"
)

func main() {
	staticDirectory := os.Getenv("STATIC_DIR")
	http.ListenAndServe(":8080", http.FileServer(http.Dir(staticDirectory)))
}
