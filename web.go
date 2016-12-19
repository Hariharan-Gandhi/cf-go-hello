package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi from GO Server")
}