package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "banana", "orange"}
	n := "pear"

	fm := make(map[string]int)
	for i, v := range fruits {
		fm[v] = i
	}
	if _, ok := fm[n]; ok {
		fmt.Println("in")
	} else {
		fmt.Println("not in")
	}

	//对 slice 进行升序排序
	ss := sort.StringSlice(fruits)
	ss.Sort()
	fruitsSorted := []string(ss)

	//查找字符串
	pos := sort.SearchStrings(fruitsSorted, n)
	if pos != len(fruitsSorted) {
		if n == fruitsSorted[pos] {
			fmt.Println("in")
		} else {
			fmt.Println("not in")
		}
	} else {
		fmt.Println("not in")
	}
}
