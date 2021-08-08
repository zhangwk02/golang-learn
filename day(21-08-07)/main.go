package main

import (
	"fmt"
)

/* func defferdemo() {

	fmt.Println("start")
	defer fmt.Println("111") // defer延时执行语句，函数即将返回的时候开始执行，按照defer顺序逆序执行
	defer fmt.Println("222")
	defer fmt.Println("333")
	fmt.Println("end")
}
*/

// defer
func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x // 给返回值赋值，然后再返回，如果有defer在这之间执行
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值是x ,将5赋值给x ,然后执行 x++ x值变成了6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 修改x
	}()
	return x // y=x=5,y才是返回值
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变还是中的副本
	}(x) // 这里x是形参数，不改变外面x的值
	return 5 // 返回值 = x = 5
}

func main() {

	fmt.Println("Helllo World")
	// defferdemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	x := 5
	y := x // 这里是值传递，单独拷贝一份
	x++
	fmt.Println("*===================*")
	fmt.Println(x, y)
	fmt.Println(&x, &y)

}
