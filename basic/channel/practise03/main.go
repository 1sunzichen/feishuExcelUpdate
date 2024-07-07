package main

import (
	"fmt"
	"time"
)

// 两个协程交替打印10个字母和数字
func main() {
	chan1 := make(chan struct{}, 1)
	chan2 := make(chan struct{}, 1)
	chan1 <- struct{}{}
	go chan1Fn(chan1, chan2)
	go chan2Fn(chan1, chan2)
	time.Sleep(time.Second * 10)
}
func chan1Fn(ch1, ch2 chan struct{}) {
	for i := 1; i < 30; i += 2 {
		<-ch1
		fmt.Printf("%d%d ", i, i+1)
		ch2 <- struct{}{}
	}
}
func chan2Fn(ch2, ch1 chan struct{}) {
	for i := 65; i < 90; i += 2 {
		<-ch1
		fmt.Printf("%c%c ", i, i+1)
		ch2 <- struct{}{}
	}
}
