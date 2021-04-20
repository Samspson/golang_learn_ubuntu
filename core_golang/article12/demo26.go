package main

import (
	"fmt"
)

// Printer 函数签名：函数的参数列表和结果列表的统称
type Printer func(contents string) (n int, err error)

func printToSd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer
	p = printToSd
	p("something")
}
