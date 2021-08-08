package main

import "fmt"

// 闭包
func cals(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base

	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := cals(10)        // 调用函数，除了返回内部函数外，还会带着外部变量base
	fmt.Println(f1(1), f2(2)) // 外部函数参数base ,被内部函数引用；每次调用函数时候，外部函数参数base也会跟着改变
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))

}
