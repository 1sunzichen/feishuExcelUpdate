package main

import "fmt"

func lengthOfLongestSubstring(s string) (ans int) {
	lens := len(s)
	if lens < 2 {
		return lens
	}
	i, j, mp := 0, 0, [128]bool{}
	for j = 0; j < lens; j++ {
		//增加右边界 计算 不重复的 字符长度
		if mp[int(s[j])] == false {
			mp[int(s[j])] = true
		} else { //mp[int(s[j])]==true
			//增加左边界索引
			for ; i < j; i++ {
				//找到重复的左边界  移动索引 继续 增加右边界
				if s[i] == s[j] {

					// if i < j {
					// 	i++
					// }
					i++
					break
				}
				//重置字符串状态
				mp[int(s[i])] = false
			}
		}
		if j-i+1 > ans {
			ans = j - i + 1
		}
	}
	return

}
func main() {
	fmt.Println(lengthOfLongestSubstring("asneriissoascderdd"))
}
