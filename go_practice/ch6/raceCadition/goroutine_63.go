package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter是所有goroutine都要增加其值的变量
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go inCounter(1)
	go inCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func inCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 捕获counter的值
		value := counter

		// 当前goroutine从线程退出，并放回到队列
		runtime.Gosched()

		// 增加本地value变量的值
		value++

		// 将该值保存会counter
		counter = value
	}
}
