package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 方法使用值接受者
/* func (c cat) move() {
	fmt.Println("猫步。。。")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
} */

// 使用指针接受者，实现方法
func (c *cat) move() {
	fmt.Println("猫步。。。")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 animal
	c1 := cat{"tom", 4}
	c2 := &cat{"咋咋", 4}
	a1 = &c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

}
