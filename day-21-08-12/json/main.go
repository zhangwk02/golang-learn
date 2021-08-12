package main

import (
	"encoding/json"
	"fmt"
)

// 序列化：将go语言的结构体变量 --> json格式的字符串
// 反序列化：将json格式的字符串 --> go语言中的结构体变量
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		"茄子",
		100,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("序列化失败，err: %v\n", err)
		return
	}
	fmt.Printf("%v\n", string(b))
	// 反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针，为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)

}
