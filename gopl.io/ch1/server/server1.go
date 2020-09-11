package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
}
