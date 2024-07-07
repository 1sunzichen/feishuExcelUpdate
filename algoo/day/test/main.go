package main

import (
	"container/list"
	"fmt"
)

func dfs(nums []int, target, i, j int) int {
	if i > j {
		return -1
	}
	m := i + (j-i)>>1
	if nums[m] > target {
		return dfs(nums, target, 0, m-1)
	} else if nums[m] < target {
		return dfs(nums, target, m+1, j)
	} else {
		return m
	}
}
func main() {
	// nums := []int{1, 2, 4, 5, 6, 7}
	// fmt.Println(dfs(nums, 5, 0, len(nums)-1))

	A := list.New()
	A.PushBack(1)
	A.PushBack(2)
	A.PushBack(3)
	A.PushBack(4)
	B := list.New()
	C := list.New()
	HannuotaSrc(A, B, C)
	for e := A.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, "A")
	}
	for e := B.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, "B")
	}
	for e := C.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, "C")
	}
}

func move(src, tar *list.List) {
	pan := src.Back()
	tar.PushBack(pan.Value)
	src.Remove(pan)
}
func Hannuota(i int, A, B, C *list.List) {
	if i == 1 {
		move(A, C)
		return
	}
	Hannuota(i-1, A, C, B)
	move(A, C)
	Hannuota(i-1, B, A, C)
}
func HannuotaSrc(A, B, C *list.List) {
	n := A.Len()
	Hannuota(n, A, B, C)
}
