package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(rw http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	defer mu.Unlock()
	fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
}

func counter(rw http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(rw, "Count = %d\n", count)
	defer mu.Unlock()
}
