package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)

func test() {

	a := <-ch1
	fmt.Println(a)
}

func main() {
	go test()

	ch1 <- 4
	// go rangech()
	// runtime.Gosched()
	time.Sleep(1 * time.Second)

}
