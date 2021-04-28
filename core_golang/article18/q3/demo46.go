package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Errno int

func (e Errno) Error() string {
	return "errno" + strconv.Itoa(int(e))
}

func main() {
	var err error

	// Example 1
	_, err = exec.LookPath(os.DevNull)
	fmt.Printf("error: %s\n", err)
	if execErr, ok := err.(*exec.Error); ok {
		execErr.Name = os.TempDir()
		execErr.Err = os.ErrNotExist
	}
	fmt.Printf("error: %s\n", err)
	fmt.Println()

	// Example 2
	err = os.ErrPermission
	if os.IsPermission(err) {
		fmt.Printf("error(permission): %s\n", err)
	} else {
		fmt.Printf("error(other): %s\n", err)
	}
	os.ErrPermission = os.ErrExist

	if os.IsPermission(err) {
		fmt.Printf("error(permission): %s\n", err)
	} else {
		fmt.Printf("error(other): %s\n", err)
	}
	fmt.Println()

	// Example 3
	const (
		ERR0 = Errno(0)
		ERR1 = Errno(1)
		ERR2 = Errno(2)
	)
	var myErr error = Errno(0)
	switch myErr {
	case ERR0:
		fmt.Println("ERR0")
	case ERR1:
		fmt.Println("ERR1")
	case ERR2:
		fmt.Println("ERR2")
	}
}
