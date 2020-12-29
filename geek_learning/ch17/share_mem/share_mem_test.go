package sharemem

import (
	"sync"
	"testing"
)

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	// time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}
