package main

import (
	"fmt"
	_ "net/http/pprof"
)

func main() {
	// nums := []int{0, 1, 2, 7, 12, 15}
	nums := []int{1, 2, 1, 2, 1}
	// target := 22
	target := 3
	fmt.Println(subarraySum(nums, target))
}

func testSums(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		if n, ok := m[target-nums[i]]; ok {
			return []int{n, i}
		}

		m[nums[i]] = i

	}
	return nil

}
func subarraySum(nums []int, k int) (index int) {
	m := make([]int, 128)
	m[0] = nums[0]
	if m[0] == k {
		return 1
	}

	for i := 1; i < len(nums); i++ {

		m[i] = nums[i] + m[i-1]
		if m[i] == k {
			index = i + 1
			break
		} else if m[i] > k {
			index = -1
			break
		}
	}
	return
}
