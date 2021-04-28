package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Printf("unexpected error: %s\n", err)
		return
	}

	r.Close()
	_, err = w.Write([]byte("hi"))
	uError := underlyingError(err)
	fmt.Printf("underlying error: %s (type: %T)\n", uError, uError)
	fmt.Println()

	paths := []string{
		os.Args[0],           // 当前的源码文件或可执行文件
		"/it/must/not/exist", //肯定不存在的目录
		os.DevNull,           // 肯定存在的目录
	}
	printError := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}
		err = underlyingError(err)
		switch err {
		case os.ErrClosed:
			fmt.Printf("err(closed)[%d]: %s\n", i, err)
		case os.ErrInvalid:
			fmt.Printf("err(invalid)[%d]: %s\n", i, err)
		case os.ErrPermission:
			fmt.Printf("err(permission)[%d]: %s\n", i, err)
		}
	}
	var f *os.File
	var index int
	{
		index = 0
		f, err = os.Open(paths[index])
		if err != nil {
			fmt.Printf("unexpected error: %s\n", err)
			return
		}

		f.Close()
		_, err := f.Read([]byte{})
		printError(index, err)
	}
	{
		index = 1
		f, _ := os.Open(paths[index])
		_, err = f.Stat()
		printError(index, err)
	}
	{
		index = 2
		_, err = exec.LookPath(paths[index])
		printError(index, err)
	}
	if f != nil {
		f.Close()
	}
	fmt.Println()

	paths2 := []string{
		runtime.GOROOT(),     // 当前环境下的GO语言根目录
		"/it/must/not/exist", //肯定不存在的目录
		os.DevNull,           // 肯定存在的目录
	}
	printError2 := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}
		err = underlyingError(err)
		if os.IsExist(err) {
			fmt.Printf("error(exist)[%d]: %s\n", i, err)
		} else if os.IsNotExist(err) {
			fmt.Printf("error(not exist)[%d]: %s\n", i, err)
		} else if os.IsPermission(err) {
			fmt.Printf("error(permission)[%d]: %s\n", i, err)
		} else {
			fmt.Printf("error(other)[%d]: %s\n", i, err)
		}
	}
	{
		index = 0
		err = os.Mkdir(paths2[index], 0700)
		printError2(index, err)
	}
	{
		index = 1
		f, err = os.Open(paths2[index])
		printError2(index, err)
	}
	{
		index = 2
		_, err = exec.LookPath(paths2[index])
	}
	if f != nil {
		f.Close()
	}
}
