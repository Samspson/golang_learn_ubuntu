package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func main() {
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
	fmt.Println()

	err1 := fmt.Errorf("invalid contents: %s", "#$%")
	err2 := errors.New(fmt.Sprintf("invalid contents: %s", "#$%"))
	if err1.Error() == err2.Error() {
		fmt.Println("The error message in err1 and err2 are the same.")
	}
}
