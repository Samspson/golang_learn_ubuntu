package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go func() {
			fmt.Println(i)
			// wg.Done()
		}()

	}
	// wg.Wait()
	time.Sleep(time.Millisecond * 500)
}
