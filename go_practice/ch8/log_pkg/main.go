package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println
	log.Println("message")

	log.Fatalln("fatal message")

	log.Panicln("panic message")
}
