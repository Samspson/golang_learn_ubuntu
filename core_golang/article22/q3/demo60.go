package main

import (
	"log"
	"sync"
	"time"
)

type counter struct {
	num uint // 计数
	mu  sync.Mutex
}

func (c *counter) number() uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.num
}

func (c *counter) add(increment uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += increment
	return c.num
}

func main() {
	c := counter{}
	count(&c)
	redundantUnlock()
}

func count(c *counter) {
	sign := make(chan struct{}, 3)
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for j := 1; j <= 20; j++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("The number is counter: %d [%d-%d]", c.number(), 1, j)
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for k := 0; k <= 20; k++ {
			time.Sleep(time.Millisecond * 300)
			log.Printf("The number is counter: %d [%d-%d]", c.number(), 2, k)
		}
	}()
	<-sign
	<-sign
	<-sign
}

func redundantUnlock() {
	var rwMu sync.RWMutex

	// rwMu.Unlock() // 引发panic

	// rwMu.RUnlock() // 引发panic

	rwMu.RLock()
	// rwMu.Unlock() // 引发panic
	rwMu.RUnlock()

	rwMu.Lock()
	// rwMu.RUnlock() // 引发panic
	rwMu.Unlock()
}
