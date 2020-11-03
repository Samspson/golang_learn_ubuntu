package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.Has(x) {
		s.words[word] ^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = append([]uint64{})
}

func (s *IntSet) Copy() *IntSet {
	intSet := &IntSet{
		words: []uint64{},
	}
	for _, value := range s.words {
		intSet.words = append(intSet.words, value)
	}
	return intSet
}

func (s *IntSet) AddAll(args ...int) {
	for _, x := range args {
		s.Add(x)
	}
}

//A与B的并集，A与B中均出现
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			continue
		}
		s.words[i] &= tword
	}
}

//A与B的差集，元素出现在A未出现在B
func (s *IntSet) DifferenceWith(t *IntSet) {
	t1 := t.Copy() //为了不改变传参t，拷贝一份
	t1.IntersectWith(s)
	for i, tword := range t1.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}
}

//A与B的并差集，元素出现在A没有出现在B，或出现在B没有出现在A
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//获取比特数组中的所有元素的slice集合
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println("x:", x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println("y:", y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println("x unionWith y:", x.String())         // "{1 9 42 144}"
	fmt.Println("x has 9,123:", x.Has(9), x.Has(123)) // "true false"
	fmt.Println("x len:", x.Len())                    //4
	fmt.Println("y len:", y.Len())                    //2
	x.Remove(42)
	fmt.Println("x after Remove 42:", x.String()) //{1 9 144}
	z := x.Copy()
	fmt.Println("z copy from x:", z.String()) //{1 9 144}
	x.Clear()
	fmt.Println("clear x:", x.String()) //{}
	x.AddAll(1, 2, 9)
	fmt.Println("x addAll 1,2,9:", x.String()) //{1 2 9}
	x.IntersectWith(&y)
	fmt.Println("x intersectWith y:", x.String()) //{9}
	x.AddAll(1, 2)
	fmt.Println("x addAll 1,2:", x.String()) //{1 2 9}
	x.DifferenceWith(&y)
	fmt.Println("x differenceWith y:", x.String()) //{1 2}
	x.AddAll(9, 144)
	fmt.Println("x addAll 9,144:", x.String()) //{1 2 9 144}
	x.SymmetricDifference(&y)
	fmt.Println("x symmetricDifference y:", x.String()) //{1 2 42 144}
	for _, value := range x.Elems() {
		fmt.Print(value, " ") //1 2 42 144
	}
}
