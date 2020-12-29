package main

import (
	"fmt"
)

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

/* result = 1
func f1() (result int) {
	result = 0
	func() {
		result++
	}()
	return
}
*/

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

/* r = 5
func f2() (r int) {
	t := 5
	r = t
	func() {
		t = t + 5
	}()
	return
}
*/

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

/* r=1
func f3() (r int) {
	r = 1
	func(r int) {
		r = r + 5
	}(r)
	return
}
*/

func main() {
	fmt.Println(f1()) // 1
	fmt.Println(f2()) // 5
	fmt.Println(f3()) // 1
}
