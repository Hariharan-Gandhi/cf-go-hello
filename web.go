package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	// http.ListenAndServe(":8090", nil)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	i := time.Now().UTC().Format(time.RFC3339Nano)
	fmt.Printf("Ping - %s\n", i)
	io.WriteString(w, fmt.Sprintf("Pong - %s\n", i))
}
