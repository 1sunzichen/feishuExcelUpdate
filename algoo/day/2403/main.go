package main

import (
	"container/list"
	"fmt"
)

func move(src, tar *list.List) {
	pan := src.Back()
	tar.PushBack(pan.Value)
	src.Remove(pan)
}
func Hanmuota(i int, A, B, C *list.List) {
	if i == 1 {
		move(A, C)
		return
	}
	Hanmuota(i-1, A, C, B)
	move(A, C)
	Hanmuota(i-1, B, A, C)

}
func solveHannuota(A, B, C *list.List) {
	n := A.Len()
	Hanmuota(n, A, B, C)
}
func main() {
	a := list.New()
	a.PushBack(1)
	a.PushBack(2)
	a.PushBack(3)
	a.PushBack(4)
	b := list.New()
	c := list.New()
	solveHannuota(a, b, c)
	for e := a.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, "A")
	}
	for e := b.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, "B")
	}
	for e := c.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, "C")
	}
}
