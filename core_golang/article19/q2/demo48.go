package main

import (
	"fmt"
	"strconv"
)

func main() {
	cid1, err := strconv.ParseInt("00001", 10, 32)
	if err == nil {
		fmt.Printf("%d, %T\n", int32(cid1), int32(cid1))
	}
	cid2, err := strconv.Atoi("00001")
	if err == nil {
		fmt.Printf("%d, %T\n", cid2, cid2)
	}

}
