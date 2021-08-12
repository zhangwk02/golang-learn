package main

import "fmt"

// 结构体模拟实现其他语言继承

type animal struct {
	name string
}

// 给animal 实现一个移动方法
func (a animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type dog struct {
	feet   uint8
	animal //animal 有的方法，dog也有
}

func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{"和尚"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
