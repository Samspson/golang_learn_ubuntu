package panicrecover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recoverd from ", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("something wrong!"))
	// os.Exit(1)
}
