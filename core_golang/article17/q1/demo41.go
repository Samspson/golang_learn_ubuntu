package main

import (
	"fmt"
	"reflect"
)

func main() {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reflect.TypeOf(numbers1))
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println("numbers1:", numbers1)

	//
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(reflect.TypeOf(numbers2))
	maxIndex2 := len(numbers2) - 1
	for i, v := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += v
		} else {
			numbers2[i+1] += v
		}
	}
	fmt.Println("numbers2:", numbers2)
}
