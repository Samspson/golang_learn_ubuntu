package main

import "fmt"

func main() {
	for index := 0; index < 5; index++ {
		defer func(index *int) { fmt.Printf("%d ", *index) }(&index)
	}
}
