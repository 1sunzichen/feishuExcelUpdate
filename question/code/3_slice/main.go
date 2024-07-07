package main

import "fmt"

func SliceCap() {
	var array [10]int
	var slice = array[5:6]
	fmt.Printf("len(slice)=%d\n", len(slice))
	fmt.Printf("cap(slice)=%d\n", cap(slice))
	//1,5
}

func main() {
	// SliceCap()
	// SlicePrint()
	// SliceExtend()
	a := []int{}
	// a = append(a, 3)
	SliceTest(a)
}

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}
func SlicePrint() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	SliceRise(s1)
	SliceRise(s2)
	fmt.Println(s1, s2)
	//1,2 2,3,4
}

// append 一个元素， 都会是原来容量的2倍，小于1024，大于1024 1.25倍 。1.18之后 小于256 为2倍。大于256
// (size+smallSizeDiv-1)/smallSizeDiv = 5
// newcap = int(capmem / ptrSize) // 6
// 两个魔法数组
// append 多个元素 ，第一个元素 原来的长度大于等于1时增加为原来容量的2倍，长度为0容量就变成1 。
// 容量4之前每次1，后续长度32之前，容量每次增加2, 64之前每次4，96之前每次8, 192每次16  ...
func SliceExtend() {
	// var slice []int
	slice := []int{}
	// slice := make([]int, 0)
	s3 := append(slice, 1, 2, 3, 5, 6, 7)
	s4 := append(s3, 3)
	//1 4  2
	//2 4  4
	//3 6  4
	s5 := append(s3, 4)
	fmt.Println(&s3[0] == &s5[0], cap(s3), cap(s5), len(slice), cap(slice), s3, s5)
	fmt.Printf("s4 cap: %d", cap(s4))
}

func SliceTest(slice []int) {
	var currentslice []int

	slice = make([]int, len(slice)+1)

	currentslice = slice

	currentslice = append([]int{}, currentslice...)
	fmt.Printf(" len%d,cap %d", len(currentslice), cap(currentslice))
	fmt.Println()
	if len(currentslice) < 1024 {
		SliceTest(currentslice)
	}
}
