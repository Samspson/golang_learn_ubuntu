package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// IntStrMap ...
type IntStrMap struct {
	m sync.Map
}

// Delete ....
func (iMap *IntStrMap) Delete(key int) {
	iMap.m.Delete(key)
}

// Load ...
func (iMap *IntStrMap) Load(key int) (value string, ok bool) {
	v, ok := iMap.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

// LoadOrStore ...
func (iMap *IntStrMap) LoadOrStore(key int, value string) (actual string, loaded bool) {
	a, loaded := iMap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

// Range ...
func (iMap *IntStrMap) Range(f func(key int, value string) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int), value.(string))
	}
	iMap.m.Range(f1)
}

// Store ...
func (iMap *IntStrMap) Store(key int, value string) {
	iMap.m.Store(key, value)
}

// ConcurrentMap 代表可自定义键类型和值类型的并发安全字典
type ConcurrentMap struct {
	m         sync.Map
	keyType   reflect.Type
	valueType reflect.Type
}

// NewConcurrentMap ...
func NewConcurrentMap(keyType, valueType reflect.Type) (*ConcurrentMap, error) {
	if keyType == nil {
		return nil, errors.New("nil key type")
	}

	if !keyType.Comparable() {
		return nil, fmt.Errorf("incomparable key type: %s", keyType)
	}

	if valueType == nil {
		return nil, errors.New("nil value type")
	}

	cMap := &ConcurrentMap{
		keyType:   keyType,
		valueType: valueType,
	}
	return cMap, nil
}

// Delete ....
func (cMap *ConcurrentMap) Delete(key interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}
	cMap.m.Delete(key)
}

// Load ...
func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}
	return cMap.m.Load(key)
}

// LoadOrStore ...
func (cMap *ConcurrentMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}
	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}
	actual, loaded = cMap.m.LoadOrStore(key, value)
	return
}

// Range ...
func (cMap *ConcurrentMap) Range(f func(key, value interface{}) bool) {
	cMap.m.Range(f)
}

// Store ...
func (cMap *ConcurrentMap) Store(key, value interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}
	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}
	cMap.m.Store(key, value)
}

var pairs = []struct {
	k int
	v string
}{
	{k: 1, v: "a"},
	{k: 2, v: "b"},
	{k: 3, v: "c"},
	{k: 4, v: "d"},
}

func main() {
	var sMap sync.Map
	// sMap.Store([]int{1, 2, 3, 4}) // 引发panic
	_ = sMap
	{
		var iMap IntStrMap
		iMap.Store(pairs[0].k, pairs[0].v)
		iMap.Store(pairs[1].k, pairs[1].v)
		iMap.Store(pairs[2].k, pairs[2].v)
		fmt.Println("[Three pairs have been stored int the ConcurrentMap instance]")

		iMap.Range(func(key int, value string) bool {
			fmt.Printf("The result of an iteration in Range: %v, %v\n", key, value)
			return true
		})

		k0 := pairs[0].k
		v0, ok := iMap.Load(k0)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n", v0, ok, k0)

		k3 := pairs[3].k
		v3, ok := iMap.Load(k3)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n", v3, ok, k3)

		k2, v2 := pairs[2].k, pairs[2].v
		actual2, loaded2 := iMap.LoadOrStore(k2, v2)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n", actual2, loaded2, k2, v2)
		v3 = pairs[3].v
		actual3, loaded3 := iMap.LoadOrStore(k3, v3)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n", actual3, loaded3, k3, v3)

		k1 := pairs[1].k
		iMap.Delete(k1)
		fmt.Printf("[The pairs with the key of %v has been removed form the ConcurrentMap instance]\n", k1)
		v1, ok := iMap.Load(k1)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n", v1, ok, k1)
		v1 = pairs[1].v
		actual1, loaded1 := iMap.LoadOrStore(k1, v1)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n", actual1, loaded1, k1, v1)

		iMap.Range(func(key int, value string) bool {
			fmt.Printf("The result of an iteration in Range: %v, %v\n", key, value)
			return true
		})
	}
	fmt.Println()
	{
		cMap, err := NewConcurrentMap(reflect.TypeOf(pairs[0].k), reflect.TypeOf(pairs[0].v))
		if err != nil {
			fmt.Printf("fatal error: %s", err)
			return
		}
		cMap.Store(pairs[0].k, pairs[0].v)
		cMap.Store(pairs[1].k, pairs[1].v)
		cMap.Store(pairs[2].k, pairs[2].v)
		fmt.Println("[Three pairs have been stored int the ConcurrentMap instance]")

		cMap.Range(func(key, value interface{}) bool {
			fmt.Printf("The result of an iteration in Range: %v, %v\n", key, value)
			return true
		})

		k0 := pairs[0].k
		v0, ok := cMap.Load(k0)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n", v0, ok, k0)

		k3 := pairs[3].k
		v3, ok := cMap.Load(k3)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n", v3, ok, k3)

		k2, v2 := pairs[2].k, pairs[2].v
		actual2, loaded2 := cMap.LoadOrStore(k2, v2)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n", actual2, loaded2, k2, v2)
		v3 = pairs[3].v
		actual3, loaded3 := cMap.LoadOrStore(k3, v3)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n", actual3, loaded3, k3, v3)

		k1 := pairs[1].k
		cMap.Delete(k1)
		fmt.Printf("[The pairs with the key of %v has been removed form the ConcurrentMap instance]\n", k1)
		v1, ok := cMap.Load(k1)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n", v1, ok, k1)
		v1 = pairs[1].v
		actual1, loaded1 := cMap.LoadOrStore(k1, v1)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n", actual1, loaded1, k1, v1)

		cMap.Range(func(key, value interface{}) bool {
			fmt.Printf("The result of an iteration in Range: %v, %v\n", key, value)
			return true
		})
	}
}
