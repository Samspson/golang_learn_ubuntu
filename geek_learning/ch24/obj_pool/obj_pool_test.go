package objpol

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// if err := pool.Release(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.Release(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
