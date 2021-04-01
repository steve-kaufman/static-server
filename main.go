package main

import (
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(os.Args[1])))
}
