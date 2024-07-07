package main

import (
	"fmt"
	"math"
	"time"
)

func do(i int, ch chan struct{}) {

	fmt.Println(i)
	time.Sleep(time.Second)
	<-ch
}

func main() {
	c := make(chan struct{}, 3000)
	for i := 0; i < math.MaxInt64; i++ {
		c <- struct{}{}
		go do(i, c)
	}
}
