package main

import "fmt"

// https://leetcode.cn/problems/partition-array-into-three-parts-with-equal-sum/solutions/616859/golangqian-zhui-he-shuang-bai-by-flash50-96u8/
func main() {
	// var arr = []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	// var arr2 = []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	// var arr3 = []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	var arr4 = []int{10, -10, 10, -10, 10, -10, 10, -10}
	// fmt.Println(canThreePartsEqualSum(arr))
	// fmt.Println(canThreePartsEqualSum(arr2))
	// fmt.Println(canThreePartsEqualSum(arr3))
	fmt.Println(canThreePartsEqualSum(arr4))
}

// 前缀和 累加
// 后缀和 累加 for循环 大于等于0  继续相加
// 遍历 前缀和 没达到 三倍 的时候，继续。
// 遍历 后端和 是否等于前缀和 返回true 三等份 从 索引2开始
func canThreePartsEqualSum(arr []int) bool {
	N := len(arr)
	presum := make([]int, N)
	presum[0] = arr[0]

	for i := 1; i < N; i++ {
		presum[i] = presum[i-1] + arr[i]
	}
	if presum[N-1]%3 != 0 {
		return false
	}
	nextsum := make([]int, N)
	nextsum[N-1] = arr[N-1]
	for i := N - 2; i > 0; i-- {
		nextsum[i] = nextsum[i+1] + arr[i]
	}
	for i := 0; i < N; i++ {
		if presum[i]*3 != presum[N-1] {
			continue
		}
		for j := i + 2; j < N; j++ {
			if presum[i] == nextsum[j] {
				return true
			}
		}
	}
	return false
}
