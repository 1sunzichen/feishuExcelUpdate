package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	ans := 0
	for key, _ := range m {
		if !m[key-1] {
			long := 1
			for m[key+1] {
				key++
				long++
			}
			if long > ans {
				ans = long
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(longestConsecutive([]int{34, 1, 3, 4, 5, 6, 35, 36, 37, 38, 39, 33}))
}
