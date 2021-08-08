package main

import "fmt"

func adder1() func(int) int {
	x := 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret1 := adder1()

	ret2 := ret1(200)
	fmt.Println(ret2)

	ret3 := adder2(100)
	ret4 := ret3(200)
	fmt.Println(ret4)
}
