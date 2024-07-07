package main

import "fmt"

func main() {
	test()
}
func test() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
