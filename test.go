package main

import (
	"database/sql"
	"fmt"
)

type Bitset struct {
	Bits []byte
}
type Region struct {
	ID     int64
	ZoneID int64
	Mask   *Bitset
}

const MaxCameraIdxPlusOne = 128

func main() {
	// for index := 0; index < 5; index++ {
	// 	defer func(index *int) { fmt.Printf("%d ", *index) }(&index)
	// }

	regions := make([]Region, 0, 32)
	fmt.Println(regions)
	var rid sql.NullInt64
	fmt.Println(rid)

	var rmask []byte
	dmask := NewBitset(rmask, MaxCameraIdxPlusOne)
	fmt.Println(dmask)
	var zoneID = int64(1)

	r := Region{
		ID:     rid.Int64,
		ZoneID: zoneID,
		Mask:   NewBitset(rmask, MaxCameraIdxPlusOne),
	}
	regions = append(regions, r)
	fmt.Println("regions:", regions[0].Mask)
	for i := range regions {
		idx := regions[i].Mask.FindAvailable()
		fmt.Println("idx:", idx)
		regions[i].Mask.Set(idx)
	}

	fmt.Println("regions:", regions[0].Mask)

	var idxxx int
	idxxx = 2
	i := idxxx / 8
	j := uint(idxxx % 8)
	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println(byte(1 << j))
}

func NewBitset(b []byte, bits int) *Bitset {
	v := make([]byte, (bits+7)/8)
	fmt.Println("len(v):", len(v))
	copy(v, b)
	return &Bitset{
		Bits: v,
	}
}

func (b *Bitset) FindAvailable() int {
	for idx := 1; idx < len(b.Bits)*8; idx++ {
		if !b.IsSet(idx) {
			return idx
		}
	}
	return 0
}

func (b *Bitset) IsSet(idx int) bool {
	i := idx / 8
	j := uint(idx % 8)
	return (b.Bits[i] & byte(1<<j)) != 0
}

func (b *Bitset) Set(idx int) {
	i := idx / 8
	j := uint(idx % 8)
	b.Bits[i] |= byte(1 << j)
}

// var intMap map[int]int

// func main() {
// 	printMemStats("初始化")

// 	// 添加1w个map值
// 	intMap = make(map[int]int, 10000)
// 	fmt.Println(intMap)
// 	for i := 0; i < 10000; i++ {
// 		intMap[i] = i
// 	}
// 	// 手动进行gc操作
// 	runtime.GC()
// 	// 再次查看数据
// 	printMemStats("增加map数据后")

// 	log.Println("删除前数组长度：", len(intMap))
// 	for i := 0; i < 10000; i++ {
// 		delete(intMap, i)
// 	}
// 	log.Println("删除后数组长度：", len(intMap))

// 	// 再次进行手动GC回收
// 	runtime.GC()
// 	printMemStats("删除map数据后")

// 	// 设置为nil进行回收
// 	intMap = nil
// 	runtime.GC()
// 	printMemStats("设置为nil后")
// 	fmt.Println(intMap)
// 	intMap = make(map[int]int, 10000)
// 	intMap[1] = 4
// 	fmt.Println(intMap)
// }

// func printMemStats(mag string) {
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	log.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
// }
