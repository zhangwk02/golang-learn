package main

import (
	"fmt"
)

func main() {

	/* for i := 1; i < 10; i++ {
		// fmt.Println(i)
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
			//fmt.Printf(" ")
		}
		// fmt.Printf("\n")
		fmt.Println()

	}
	s0 := "这一次，等待"
	fmt.Println(len(s0))
	for i, v := range s0 {
		fmt.Printf("%d %v\n", i, v)
	}

	s1 := "Hello"
	fmt.Println(len(s1))
	for i, _ := range s1 {
		fmt.Printf("%d %c\n", i, s1[i])
	} */
	/* 	switch n := 7; n {
	   	case 1, 3, 5:
	   		fmt.Println("奇数")
	   	case 2, 4, 6:
	   		fmt.Println("偶数")
	   	default:
	   		fmt.Println("非奇非偶")
	   	}
	*/
	// 指针
	/* 	var a *int
	   	fmt.Println(a)
	*/
	// map

	/* 	var m1 map[string]int // 定义map,还未初始化即没有在内存中开辟空间
	   	fmt.Printf("%T\n", m1)
	   	m1 = make(map[string]int, 10)
	   	fmt.Println(m1 == nil)
	   	m1["数学"] = 90
	   	m1["语文"] = 70
	   	fmt.Printf("%T\n", m1)
	   	fmt.Println(m1)
	   	v, isexit := m1["古力娜扎"]
	   	if !isexit {
	   		fmt.Println(m1["古力娜扎"], "不存在")
	   	} else {
	   		fmt.Println(v)
	   	} */

	// 切片
	/* 	arry := [10]int{1, 2, 3, 5, 6}
	   	fmt.Println(arry)
	   	a1 := arry[:3]
	   	a2 := a1
	   	fmt.Println(a1)
	   	a2[0] = 200
	   	fmt.Println(a1, a2, arry)
	*/

	// 指针
	// go语言里面的指针只能读不能修改，不能修改指针变量指向的地址
	/* 	addr := "色木"
	   	addrp := &addr
	   	fmt.Println(addrp)
	   	fmt.Printf("%T\n", addrp)
	   	addrv := *addrp
	   	fmt.Println(addrv) */

	// 回文判断
	//s := "1上海自来水来自海上1"
	s := "AacbcaA1"
	// for i, v := range s {
	// 	println(i, v)
	// }
	r := make([]rune, 0, len(s))
	for _, c := range s {
		r = append(r, c)
	}
	fmt.Println(r)
	i, j := 0, len(r)-1
	for i < j {
		if r[i] != r[j] {
			fmt.Println("不是回文")
			return
		}
		i++
		j--
	}
	fmt.Print("是回文")

}
