package main

import (
	"fmt"
)

// 客户购买 X 瓶酒，就可以用 Y 个空酒瓶来兑换一瓶新酒。

// 当空瓶子能换一瓶酒的时候，我们就换一瓶酒
func numWaterBottles1(numBottles, numExchange int) int {
	if numBottles < 1 || numBottles > 100 {
		panic("err numBottles")
	}
	if numExchange < 2 || numBottles > 100 {
		panic("err numExchange")
	}
	total := numBottles
	for numBottles >= numExchange {
		numBottles -= numExchange
		total++
		numBottles++
	}
	return total
}

// 每次将所有的空瓶子全部兑换完（可兑换的最大值），然后将兑换的酒再喝完，再进行下一次兑换
func numWaterBottles2(numBottles, numExchange int) int {
	if numBottles < 1 || numBottles > 100 {
		panic("err numBottles")
	}
	if numExchange < 2 || numBottles > 100 {
		panic("err numExchange")
	}
	total := numBottles
	for numBottles >= numExchange {
		n := numBottles / numExchange
		total += n
		numBottles = numBottles%numExchange + n
	}
	return total
}

func main() {
	var x, y = 9, 3
	fmt.Printf("X=%v,Y=%v,(方法1)总共可以喝%v瓶\n", x, y, numWaterBottles1(x, y)) // 13
	fmt.Printf("X=%v,Y=%v,(方法2)总共可以喝%v瓶\n", x, y, numWaterBottles2(x, y)) // 13

	x, y = 15, 4
	fmt.Printf("X=%v,Y=%v,(方法1)总共可以喝%v瓶\n", x, y, numWaterBottles1(x, y)) // 19
	fmt.Printf("X=%v,Y=%v,(方法2)总共可以喝%v瓶\n", x, y, numWaterBottles2(x, y)) // 19
	x, y = 5, 5
	fmt.Printf("X=%v,Y=%v,(方法1)总共可以喝%v瓶\n", x, y, numWaterBottles1(x, y)) // 6
	fmt.Printf("X=%v,Y=%v,(方法2)总共可以喝%v瓶\n", x, y, numWaterBottles2(x, y)) // 19
	x, y = 2, 3
	fmt.Printf("X=%v,Y=%v,(方法1)总共可以喝%v瓶\n", x, y, numWaterBottles1(x, y)) // 2
	fmt.Printf("X=%v,Y=%v,(方法2)总共可以喝%v瓶\n", x, y, numWaterBottles2(x, y)) // 19

	// x, y = 100, 1
	// fmt.Printf("X=%v,Y=%v,总共可以喝%v\n", x, y, numWaterBottles1(x, y)) // 2
}
