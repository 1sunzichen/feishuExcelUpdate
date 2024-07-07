package main

import (
	"context"
	"fmt"
	"time"
)

func fn1(ch chan struct{}) {

	time.Sleep(time.Second * 1)
	ch <- struct{}{}
}
func fn2(ch chan struct{}) {

	time.Sleep(time.Second * 3)
	ch <- struct{}{}
}

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	go func() {
		go fn1(ch1)
		select {
		case <-ctx.Done():
			fmt.Println("fm1 超时")
			break
		case <-ch1:
			fmt.Println("fm1结束done")
		}
	}()
	go func() {
		go fn2(ch2)
		select {
		case <-ctx.Done():
			fmt.Println("fm2 超时")
			break
		case <-ch2:
			fmt.Println("fm2 结束done")

		}
	}()
	time.Sleep(time.Second * 6)
}
