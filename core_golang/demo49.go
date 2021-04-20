package main

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	num    = 987654321
	numstr = "987654321"
)

// strconv.ParseInt
func BenchmarkStrconvParseInt(b *testing.B) {
	num64 := int64(num)
	for i := 0; i < b.N; i++ {
		x, err := strconv.ParseInt(numstr, 10, 64)
		if x != num64 || err != nil {
			b.Error(err)
		}
	}
}

//
func BenchmarkStrconvFormatInt(b *testing.B) {
	num64 := int64(num)
	for i := 0; i < b.N; i++ {
		x := strconv.FormatInt(numstr, 10)
		if x != num64 {
			b.Error(err)
		}
	}
}

// strconv.Atoi
func BenchmarkStrconvAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, err := strconv.Atoi(numstr)
		if x != num || err != nil {
			b.Error(err)
		}
	}
}

// fmt.Sscan
func BenchmarkFmtSscan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x int
		n, err := fmt.Sscanf(numstr, "%d", &x)
		if n != 1 || x != num || err != nil {
			b.Error(err)
		}
	}
}
