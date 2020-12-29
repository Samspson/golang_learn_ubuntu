package main

var a string

func f() {
	print(a)
}

func main() {
	a = "hello, world"
	go f()
}
