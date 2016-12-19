package main

import (
	"net/http"
	"io"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi from GO Server")
}
