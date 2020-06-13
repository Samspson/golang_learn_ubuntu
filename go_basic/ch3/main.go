package main

import "fmt"

type Interger int

// func (a *Interger) Add(b Interger) {
// 	*a += b
// }

func (a Interger) Add(b Interger) {
	a += b
}

func main() {
	var a Interger = 1
	a.Add(2)
	fmt.Println("a =", a)
}
