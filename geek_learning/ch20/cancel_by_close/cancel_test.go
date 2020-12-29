package cancel

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh:
		return true
	default:
		return false
	}
}

func cancel1(cancelCh chan struct{}) {
	cancelCh <- struct{}{}
}

func cancel2(cancelCh chan struct{}) {
	close(cancelCh)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	cancel2(cancelChan)
	time.Sleep(time.Second * 5)
}
