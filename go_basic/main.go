package main

import "fmt"

var (
	i int32
	j int64
)

func main() {
	fmt.Println(124 << 2)
	i, j = 1, 2
	if i == 1 || j == 2 {
		fmt.Println("i and j are equal.")
	}
	str := "Hello,世界"
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}
