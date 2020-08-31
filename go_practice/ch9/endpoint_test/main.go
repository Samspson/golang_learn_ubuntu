package main

import (
	"golang_learn_ubuntu/go_practice/ch9/endpoint_test/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started ： Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
