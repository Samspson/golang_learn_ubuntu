package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag =", flag)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		num := <-ch
	// 		fmt.Println(num)
	// 	}
	// 	quit <- true
	// }()
	// fibonacci(ch, quit)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num =", num)
			case <-time.After(time.Second * 3):
				fmt.Println("timeout")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	<-quit
	fmt.Println("End")
}
