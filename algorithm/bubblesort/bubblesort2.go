package main

import (
	"fmt"
)

func BubbleSort2(arr []int) {
	count := 0
	flag := false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = true
			}
		}
		if !flag {
			fmt.Println("count", count)
			return
		} else {
			flag = false
		}
	}
}

func main() {
	arr := []int{1, 3, 5, 4, 2, 6, 7, 8}
	fmt.Printf("arr befor sort: %v\n", arr)
	BubbleSort2(arr)
	fmt.Printf("arr after sort: %v\n", arr)
}
