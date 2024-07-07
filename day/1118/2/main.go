package main

import (
	"fmt"
	"runtime"
)

func add(i *int32) {
	*i += 1
	// atomic.AddInt32(i, 1)
}
func main() {
	c := int32(0)
	for i := 0; i < 100; i++ {
		go add(&c)
	}
	runtime.Gosched()
	fmt.Println(c)
}
