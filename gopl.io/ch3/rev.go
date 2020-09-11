package main

import (
	"fmt"
)

func revers(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	revers(a[:])
	fmt.Println("revers a:", a)

	s := []int{0, 1, 2, 3, 4, 5}
	revers(s[:2])
	revers(s[2:])
	revers(s)
	fmt.Println("revers s:", s)
}
