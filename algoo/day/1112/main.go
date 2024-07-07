package main

import (
	"fmt"
)

// https://leetcode.cn/problems/binary-prefix-divisible-by-5/description/
func main() {
	Test3()
}
func test() {
	nums := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}
	arrlist := prefixesDivBy5(nums)
	fmt.Println(arrlist)
}
func prefixesDivBy5(nums []int) []bool {
	ret := make([]bool, len(nums))
	last := 0
	for i := 0; i < len(nums); i++ {
		last = last*2 + nums[i]
		last = last % 5
		if last != 0 {
			ret[i] = false
		} else {
			ret[i] = true
		}
	}
	return ret
}
