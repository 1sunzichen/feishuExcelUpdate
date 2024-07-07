package main

import (
	"fmt"
	"time"
)

var dog = make(chan struct{})
var fish = make(chan struct{})
var cat = make(chan struct{})

func main() {
	var a = 1234
	go test()
	// runtime.Gosched()
	time.Sleep(time.Second * time.Duration(a))
}
func Dog() {
	<-dog
	var a = 12344
	fmt.Println("dog", a)
	fish <- struct{}{}

}
func Fish() {
	<-fish
	fmt.Println("fish")
	cat <- struct{}{}
}
func Cat() {
	<-cat

	fmt.Println("cat")
	dog <- struct{}{}
}
func test() {

	for i := 0; i < 100; i++ {
		go Dog()
		go Fish()
		go Cat()
	}
	dog <- struct{}{}
}
