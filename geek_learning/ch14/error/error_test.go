package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var ErrLessThanTwo = errors.New("n should not be less than 2")
var ErrLargerThanHundred = errors.New("n should not be larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should be in [2,100]")
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibnacci(t *testing.T) {
	if v, err := GetFibonacci(5); err != nil {
		if err == ErrLessThanTwo {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(list)
}
