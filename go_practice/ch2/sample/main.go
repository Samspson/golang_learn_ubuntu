package main

import (
	"log"
	"os"

	_ "golang_learn_ubuntu/go_practice/ch2/sample/matchers"
	"golang_learn_ubuntu/go_practice/ch2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
