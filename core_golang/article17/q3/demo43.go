package main

import "fmt"

func main() {
	// value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// switch value3[4] {
	// case 0, 1, 2:
	// 	fmt.Println("0 or 1 or 2")
	// case 2, 3, 4: // 此条编译不通过 duplicate case 2 in switch
	// 	fmt.Println("2 or 3 or 4")
	// case 4, 5, 6: // 此条编译不通过 duplicate case 4 in switch
	// 	fmt.Println("4 or 5 or 6")
	// }

	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or 6")
	}

	// value6 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// switch t := value6.(type) { // 编译不通过
	// case uint8, uint16:
	// 	fmt.Println("uint8 or uint16")
	// case byte: // cannot type switch on non-interface value value6 (type [7]int8)
	// 	fmt.Println("byte")
	// default:
	// 	fmt.Println("unsupported type: %T", t)
	// }
}
