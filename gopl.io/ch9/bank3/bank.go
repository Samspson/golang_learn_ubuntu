package bank3

import (
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

// Balance 并发安全的
func Balance() int {
	mu.Lock()
	b := balance
	// Unlock在return语句之后执行
	mu.Unlock()
	return b
}
