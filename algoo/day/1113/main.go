package main

import (
	"fmt"
	"math"
)

func main() {
	test2()
}

// https://leetcode.cn/submissions/detail/481725514/
func test(n int) string {
	var ans string
	for n > 0 {
		n--
		ans = string(n%26+'A') + ans //个位
		n /= 26                      //下一位
	}
	return ans
}

// https://leetcode.cn/problems/minimum-index-sum-of-two-lists/solutions/1335488/by-kyushu-rm13/
func test2() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Shogun", "Burger King"}
	fmt.Println(findRestaurant2(list1, list2))
}
func findRestaurant2(list1 []string, list2 []string) (ret []string) {
	cache, minIdx := map[string]int{}, math.MaxInt32
	for i := 0; i < len(list1); i++ {
		cache[list1[i]] = i
	}
	for i := 0; i < len(list2); i++ {
		j, ok := cache[list2[i]]
		if ok && i+j <= minIdx {
			if j+i < minIdx {
				ret = []string{}
				minIdx = i + j
			}
			ret = append(ret, list2[i])
		}
	}
	return
}
