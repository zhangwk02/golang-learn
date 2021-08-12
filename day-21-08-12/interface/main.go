package main

import "fmt"

// 接口，定义一个能叫的类型
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Printf("喵喵喵\n")
}

func (d dog) speak() {
	fmt.Printf("汪汪汪\n")
}

func (p person) speak() {
	fmt.Printf("啊啊啊\n")
}

func da(x speaker) {
	x.speak()
}

func main() {
	var (
		c1 cat
		d1 dog
		p1 person
	)
	da(c1)
	da(d1)
	da(p1)
	var ss speaker
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}
