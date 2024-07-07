package main

import "fmt"

func StringExaml() {
	var s string
	// s = "中国"
	fmt.Println(len(s))
}
func main() {
	StringExaml()
}

//为什么不允许修改字符串   string 内存指针
//string 和 []byte 如何取舍  p59
