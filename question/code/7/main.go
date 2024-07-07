package main

import (
	"fmt"
	"time"
)

// Q优化性能
func FindMonkey(s []string) bool {
	// A 切片 赋值内存拷贝，使用切片下标
	// for i:=range s
	// s[i]=="monkey"
	for _, v := range s {
		if v == "monkey" {
			return true
		}
	}
	return false
}

// Q发生什么
// A没有打印 永久阻塞 //nil
func RangeNilChannel() {

	var c chan string
	for v := range c {
		fmt.Println(v)
	}
}

// Q 发生什么
// A hi 之后阻塞
func RangeTimer() {
	t := time.NewTimer(time.Second)
	for _ = range t.C {
		fmt.Println("hi")
	}
}
func main() {
	RangeChannel()
}

func RangeChannel() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "Wprld"
	time.AfterFunc(time.Microsecond, func() {
		close(c)
	})
	for e := range c {
		fmt.Printf("%s\n", e)
	}
}
