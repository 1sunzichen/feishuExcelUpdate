package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("defer main g")
	go func() {
		defer func() {
			recover()
		}()
		panic("")
		fmt.Println("end g")
	}()
	time.Sleep(time.Second)
	fmt.Println("end main g")
}
