package main

import "fmt"

// 方法

type dog struct {
	name string
}

// 构造函数

func newDog(name string) dog {
	return dog{
		name,
	}
}

//方法作用域特定类型的函数
//接受dog类型变量来调用，这里接受者是dog类型变量
//接受者表示的是调用该方法的具体类型变量，多用类型首字母表示（类似java this,python self）
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}

func main() {
	d1 := newDog("哈士奇")
	d1.wang()
}
