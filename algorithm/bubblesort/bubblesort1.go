package main

import (
	"fmt"
)

func BubbleSort1(arr []int) {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("count", count)
}

func main() {
	arr := []int{1, 3, 5, 4, 2, 6, 7, 8}
	fmt.Printf("arr befor sort: %v\n", arr)
	BubbleSort1(arr)
	fmt.Printf("arr after sort: %v\n", arr)
}
