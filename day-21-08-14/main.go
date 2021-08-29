package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "这是一条日志")
	fileObj, _ := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(fileObj, "this is a log")
}
