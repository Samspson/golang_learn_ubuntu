package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// flag.Parse()
	// lib.Hello(name)
	t1 := time.Now()
	fmt.Println(t1)
	t2, _ := ptypes.TimestampProto(t1)
	fmt.Println(t2)

	downlinkHost := "12.7.0.1:5060"

}
