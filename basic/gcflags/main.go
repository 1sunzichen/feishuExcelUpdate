package main

//go build -gcflags -S 加文件名    查看汇编指令
import "fmt"

type p struct {
	age int
}

func main() {
	var b = p{age: 11}
	var a = &b
	// var a = &p{age: 11}
	fmt.Println(a) // go:nosplit
}
