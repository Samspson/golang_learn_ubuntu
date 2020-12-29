package client

import (
	"golang_learn_ubuntu/geek_learning/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
}
