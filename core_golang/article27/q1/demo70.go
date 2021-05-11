package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

//  bufPool 代表存放数据块缓冲区的临时对象池
var bufPool sync.Pool

// Buffer 代表一个简易的数据块缓冲区的接口
type Buffer interface {
	Delimiter() byte
	Write(contents string) (err error)
	Read() (contents string, err error)
	Free()
}

// myBuffer 代表了数据块缓冲区的一种实现
type myBuffer struct {
	buf       bytes.Buffer
	delimiter byte
}

func (b *myBuffer) Delimiter() byte {
	return b.delimiter
}

func (b *myBuffer) Write(contents string) (err error) {
	if _, err = b.buf.WriteString(contents); err != nil {
		return
	}
	return b.buf.WriteByte(b.delimiter)
}

func (b *myBuffer) Read() (contents string, err error) {
	return b.buf.ReadString(b.delimiter)
}

func (b *myBuffer) Free() {
	bufPool.Put(b)
}

var delimiter = byte('\n')

func init() {
	bufPool = sync.Pool{
		New: func() interface{} {
			return &myBuffer{delimiter: delimiter}
		},
	}
}

// GetBuffer 用于获取一个数据块缓冲区
func GetBuffer() Buffer {
	return bufPool.Get().(Buffer)
}

func main() {
	buf := GetBuffer()
	defer buf.Free()
	buf.Write("A Pool is a set of temporay objects that may be individually saved and retrieved.")
	buf.Write("A Pool is safe for use by multiple goroutines simulataneously.")
	buf.Write("A Pool must not be copied after first use.")

	fmt.Println("The data blocks in buffer:")
	for {
		block, err := buf.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Errorf("unexpected error: %s", err))
		}
		fmt.Print(block)
	}
}
