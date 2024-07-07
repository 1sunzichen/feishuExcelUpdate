package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan struct{}, 1)
var chan2 = make(chan struct{}, 1)

func main() {
	go chanFn1()
	go chanFn2()
	chan1 <- struct{}{}
	time.Sleep(time.Second * 5)
}
func chanFn1() {

	for i := 1; i < 30; i += 2 {
		<-chan1
		fmt.Printf("%d%d", i, i+1)
		chan2 <- struct{}{}
	}
}
func chanFn2() {

	for i := 65; i <= 90; i += 2 {
		<-chan2
		fmt.Printf("%c%c ", i, i+1)
		chan1 <- struct{}{}
	}
}
