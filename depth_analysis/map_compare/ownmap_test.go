package main

import (
	"strconv"
	"testing"
)

func RangeSliceInt(input []int, find int) (index int) {
	for index, value := range input {
		if value == find {
			//   fmt.Println("found at",index)
			return index
		}
	}
	return -1
}
func RangeSliceIntPrint(input []int) {
	for _, _ = range input {
		continue
	}
}
func MapLookupInt(input map[int]int, find int) (key, value int) {
	if value, ok := input[find]; ok {
		//  fmt.Println("found at", find,value)
		return find, value
	}
	return 0, 0
}
func MapRangeInt(input map[int]int) {
	for _, _ = range input {
		continue
	}
}
func DirectSliceInt(input []int, index int) int {
	return input[index]
}

// 对于字符串 //
func RangeSliceString(input []string, find string) (index int) {
	for index, value := range input {
		if value == find {
			// fmt.Println("found at",index)
			return index
		}
	}
	return -1
}
func RangeSliceStringPrint(input []string) {
	for _, _ = range input {
		continue
	}
}
func MapLookupString(input map[string]string, find string) (key, value string) {
	if value, ok := input[find]; ok {
		//              fmt.Println("found at", find,value)
		return find, value
	}
	return "0", "0"
}
func MapRangeString(input map[string]string) {
	for _, _ = range input {
		continue
	}
}
func DirectSliceString(input []string, index int) string {
	return input[index]
}

func Benchmark_TimeRangeSliceInt(b *testing.B) {
	b.StopTimer()
	input := make([]int, 100000, 500000)
	for i := 0; i < 100000; i++ {
		input[i] = i + 10
	}
	b.StartTimer()
	b.N = 2000000 // 只是为了避免数百万次fmt.Println（以防你在 slicemap.go 包中进行 fmt.Println）
	for i := 0; i < b.N; i++ {
		RangeSliceInt(input, 100009) // 对于最坏情况，检查最后一个元素
	}
}
func Benchmark_TimeDirectSliceInt(b *testing.B) {
	b.StopTimer()
	input := make([]int, 100000, 500000)
	for i := 0; i < 100000; i++ {
		input[i] = i + 10
	}
	b.StartTimer()
	b.N = 2000000 // 只是为了避免数百万次 fmt.Println（以防你在 slicemap.go 包中进行 fmt.Println）
	for i := 0; i < b.N; i++ {
		DirectSliceInt(input, 99999) // 直接检查索引值。o(1)。发送索引
	}
}
func Benchmark_TimeMapLookupInt(b *testing.B) {
	b.StopTimer()
	input := make(map[int]int)
	// 扔一些值到 Map 中
	for i := 1; i <= 100000; i++ {
		input[i] = i + 10
	}
	b.StartTimer()
	b.N = 2000000 // 只是为了避免数百万次fmt.Println（以防你在 slicemap.go 包中进行 fmt.Println）
	for k := 0; k < b.N; k++ {
		MapLookupInt(input, 100000)
	}
	/*
		运行命令：
		go test -bench=Benchmark_TimeMapLookup
	*/
}
func Benchmark_TimeSliceRangeInt(b *testing.B) {
	b.StopTimer()
	input := make([]int, 5, 10)
	for i := 0; i < 5; i++ {
		input[i] = i + 10
	}
	b.StartTimer()
	b.N = 2000000 // 只是为了避免数百万次fmt.Println（以防你在 slicemap.go 包中进行 fmt.Println）
	for k := 0; k < b.N; k++ {
		RangeSliceIntPrint(input)
	}
}
func Benchmark_TimeMapRangeInt(b *testing.B) {
	b.StopTimer()
	input := make(map[int]int)
	// 扔一些值到 Map 中
	for i := 1; i <= 100000; i++ {
		input[i] = i + 10
	}
	b.StartTimer()
	b.N = 2000 // 只是为了避免数百万次fmt.Println（以防你在 slicemap.go 包中进行 fmt.Println）
	for k := 0; k < b.N; k++ {
		MapRangeInt(input)
	}
}

// 测试字符串
func Benchmark_TimeRangeSliceString(b *testing.B) {
	b.StopTimer()
	input := make([]string, 100000, 500000)
	for i := 0; i < 100000; i++ {
		input[i] = strconv.FormatInt(int64(i+10), 10)
	}
	b.StartTimer()
	b.N = 2000000 // 只是为了避免数百万次fmt.Println（以防你在 slicemap.go 包中进行 fmt.Println）
	for i := 0; i < b.N; i++ {
		RangeSliceString(input, "100009") // 对于最坏情况，检查最后一个元素
	}
}
func Benchmark_TimeDirectSliceString(b *testing.B) {
	b.StopTimer()
	input := make([]string, 100000, 500000)
	for i := 0; i < 100000; i++ {
		input[i] = strconv.FormatInt(int64(i+10), 10)
	}
	b.StartTimer()
	b.N = 2000000 // 只是为了避免数百万次fmt.Println（以防你在 slicemap.go 包中进行 fmt.Println）
	for i := 0; i < b.N; i++ {
		DirectSliceString(input, 99999) // 直接检查索引值。o(1)
	}
}
func Benchmark_TimeMapLookupString(b *testing.B) {
	b.StopTimer()
	input := make(map[string]string)
	// 扔一些值到 Map 中
	for i := 1; i <= 10000; i++ {
		input[strconv.FormatInt(int64(i), 10)] = strconv.FormatInt(int64(i+10), 10)
	}
	b.StartTimer()
	b.N = 20000 // 只是为了避免数百万次fmt.Println（以防你在 slicemap.go 包中进行 fmt.Println）
	for k := 0; k < b.N; k++ {
		MapLookupString(input, "10000")
	}
	/*
		运行：
		go test -bench=Benchmark_TimeMapLookupString
	*/
}
