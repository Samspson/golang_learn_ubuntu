package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerBody)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerBody(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%s %s %s\n", r.Method, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(rw, "Host = %q\n", r.Host)
	fmt.Fprintf(rw, "RemoterAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(rw, "Form[%q] = %q\n", k, v)
	}
}
