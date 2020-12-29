package main

import (
	"fmt"
	"os"
)

var (
	name string
	id   int
	age  int
)

func main() {
	fmt.Println("-----------------------fmt.Errorf()-----------------------------")
	// fmt.Errorf()
	// %q 对于双引号的字符串进行安全转义输出。
	name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

	fmt.Println("-----------------------fmt.Fprint()-----------------------------")
	// fmt.Fprint()
	name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")

	fmt.Println("-----------------------fmt.Fprintf()-----------------------------")
	// fmt.Fprintf()
	n, err = fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

	fmt.Println("-----------------------fmt.Fprintln()----------------------------")
	// fmt.Fprintln()
	n, err = fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")
}
