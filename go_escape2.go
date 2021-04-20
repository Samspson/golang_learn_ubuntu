package main

type User struct {
	name interface{}
	// name string
}

func main() {
	name := "WilburXu"
	MyPrintln(name)
}

func MyPrintln(one interface{}) (n int, err error) {
	var userInfo = new(User)
	userInfo.name = one // 泛型赋值，逃逸
	return
}

/*
func MyPrintln(one string) (n int, err error) {
	var userInfo = new(User)
	userInfo.name = one // 泛型赋值，逃逸
	return
}
*/
