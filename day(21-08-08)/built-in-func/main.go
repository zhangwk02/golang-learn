package main

import "fmt"

// go 内置函数
func fb() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放链接数据库")
	}()
	fmt.Println("b")
	panic("出现严重错误！！！")

}

func main() {
	fb()
}
