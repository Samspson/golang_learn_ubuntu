package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	fmt.Println("len(mySlice): ", len(mySlice))
	fmt.Println("cap(mySlice): ", cap(mySlice))

	mySlice2 := []int{8, 9, 10}
	// ... 相当于打散后传入
	fmt.Println("mySlice append array: ", append(mySlice, mySlice2...))

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{7, 8, 9}
	fmt.Println("copy slice1 into slice2: ", copy(slice2, slice1))
	fmt.Println("copy slice2 into slice1: ", copy(slice1, slice2))
}
