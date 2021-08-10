// 构造函数
package main

import "fmt"

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

func newPerson(name string, age int) *person {
	return &person{
		name,
		age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("测试1", 18)
	p2 := newPerson("测试2", 25)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1, p2)
	fmt.Printf("%v\n", p1)
	fmt.Println(p1.name)
	d1 := newDog("黑狗")
	fmt.Println(d1)
	fmt.Println(d1.name)
}
