package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// mailbox代表信箱
	var mailbox uint8

	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 5
	// 发信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 0; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)

	// 收信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for j := 0; j <= max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", j)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
