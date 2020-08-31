package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter是所有goroutine都要增加其值的变量
	counter int64
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
		// 安全地对counter加1
		atomic.AddInt64(&counter, 1)

		// 当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
	}
}
