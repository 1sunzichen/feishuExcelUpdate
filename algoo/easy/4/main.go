// 746. 使用最小花费爬楼梯
// https://leetcode.cn/problems/min-cost-climbing-stairs/
package main

import "fmt"

func main() {
	a := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	b := minCostClimbingStairs(a)
	fmt.Println(b)
}
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-2], dp[n-1])
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
