package main

import "fmt"

func main() {
	nums := []int{-1, -1, -1, 0 - 1, -1}
	fmt.Println(pivotIndex(nums))
}
func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	leftnum := 0
	res := -1
	for k, v := range nums {
		// fmt.Println(float64(sum-v) / 2)
		if float64(sum-v)/2 == float64(leftnum) {
			res = k
			break
		}
		leftnum += v
	}
	return res
}
