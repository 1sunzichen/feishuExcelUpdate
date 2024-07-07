package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan struct{}, 1)
var ch2 = make(chan struct{}, 1)

func func1() {
	for i := 0; i < 10; i += 2 {
		<-ch1
		fmt.Println(i, "ch1")
		ch2 <- struct{}{}
	}
}
func func2() {
	for i := 1; i < 10; i += 2 {
		<-ch2
		fmt.Println(i, "ch2")
		ch1 <- struct{}{}
	}
}
func main() {
	ch1 <- struct{}{}
	go func1()
	go func2()
	time.Sleep(time.Second * 100)
}
