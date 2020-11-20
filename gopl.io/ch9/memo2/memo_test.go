package memo_test

import (
	"testing"

	memo "golang_learn_ubuntu/gopl.io/ch9/memo2"

	"golang_learn_ubuntu/gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
