package main

import (
	"fmt"
)

func main() {
	var j int = 5

	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i,j: %d, %d\n", i, j)
		}
	}()
	// 变量a指向的闭包函数引用了局部变量i, j, 只有内部的匿名函数才能访问变量i3.
	a()

	j *= 2

	a()
}
