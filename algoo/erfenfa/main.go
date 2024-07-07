package main

import (
	"fmt"
	"sort"
)

// 09090909090
// 99999999999
// 12345678901
func binarySearch(scores []int, target int) int {
	min := 0
	max := len(scores) - 1
	for min <= max {
		mid := min + (max-min)/2
		if scores[mid] == target {
			return mid
		} else if scores[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}
func main() {
	scores := []int{1, 3, 24, 22, 12, 32, 23, 11, 33, 26}
	sort.Ints(scores)

	index := binarySearch2(scores, 11)
	if index != -1 {
		fmt.Printf("找到了分数为11的学生，索引位置%d\n", index)

	} else {
		fmt.Println("没有找到分数为11的学生")
	}
}

func dfs(nums []int, target, i, j int) int {
	if i > j {
		return -1
	}
	m := i + ((j - i) >> 1)
	if nums[m] < target {
		return dfs(nums, target, m+1, j)
	} else if nums[m] > target {
		return dfs(nums, target, i, m-1)
	} else {
		return m
	}
}
func binarySearch2(nums []int, target int) int {
	n := len(nums)
	return dfs(nums, target, 0, n-1)
}
