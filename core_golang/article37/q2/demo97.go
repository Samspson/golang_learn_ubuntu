package main

import (
	"errors"
	"fmt"
	"golang_learn_ubuntu/core_golang/article37/common"
	"golang_learn_ubuntu/core_golang/article37/common/op"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	profileName    = "memprofile.out"
	memProfileRate = 8
)

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("memory profile creation error: %v\n", err)
		return
	}
	defer f.Close()
	startMemProfile()
	if err = common.Execute(op.MemProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
	if err := stopMemProfile(f); err != nil {
		fmt.Printf("memory profile stop error: %v\n", err)
		return
	}
}

func startMemProfile() {
	runtime.MemProfileRate = memProfileRate
}

func stopMemProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.WriteHeapProfile(f)
}
