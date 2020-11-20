package main

import (
	"fmt"
)

// 定义一个Robot结构体
type Robot struct {
	Name  string
	Color string
	Model string
}

func main() {

	fmt.Println("浅拷贝 内容和内存地址一样，改变其中一个对象的值时，另一个同时变化。")
	robot1 := Robot{
		Name:  "小白-X型-V1.0",
		Color: "白色",
		Model: "小型",
	}
	robot2 := &robot1
	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, robot2)

	fmt.Println("在这里面修改Robot1的Name和Color属性")
	robot1.Name = "小黑-X型-V1.1"
	robot1.Color = "黑色"

	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, robot2)

}
