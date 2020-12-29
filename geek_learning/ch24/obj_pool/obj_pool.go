package objpol

import (
	"errors"
	"time"
)

type ReusableObj struct{}

type ObjPool struct {
	bufChan chan *ReusableObj //用于缓冲可重用对象
}

// NewObjPool ...
func NewObjPool(numOfObj int) *ObjPool {
	objpool := ObjPool{}
	objpool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objpool.bufChan <- &ReusableObj{}
	}

	return &objpool
}

// GetObj ...
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time out")
	}
}

// Release ...
func (p *ObjPool) Release(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
