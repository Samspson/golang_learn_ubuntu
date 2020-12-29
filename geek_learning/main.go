package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 2000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			a := 0
// 			for i := 0; i < 1e6; i++ {
// 				a++
// 			}
// 			wg.Done()
// 		}()
// 		time.Sleep(100 * time.Millisecond)
// 	}
// 	wg.Wait()
// }

type Stsr struct {
	st int
}

func main() {
	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	fmt.Println("Hello trace")
}
