// 下一个更大元素 I
package main

import "fmt"

// https://leetcode.cn/problems/next-greater-element-i/solutions/1065517/xia-yi-ge-geng-da-yuan-su-i-by-leetcode-bfcoj/
func main() {
	n := []int{4, 2, 3}
	n2 := []int{2, 5, 4, 3}
	fmt.Println(nextGreaterElement(n, n2))
}

// nums1 [4,2,3] nums2[2,5,4,3]
// -1 5 -1
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, m)
	for i := 0; i < m; i++ {
		j := 0
		for j < n && nums2[j] != nums1[i] {
			j++
		}
		k := j + 1
		for k < n && nums2[k] < nums1[i] {
			k++
		}
		if k < n {
			res[i] = nums2[k]
		} else {
			res[i] = -1
		}
	}
	return res
}
