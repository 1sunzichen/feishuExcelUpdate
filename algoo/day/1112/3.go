package main

import "fmt"

func Test3() {
	s := "()()"
	str := removeOuterParentheses(s)
	fmt.Println(str, "str", s)
}

// 双指针
func removeOuterParentheses(s string) string {
	ans, n := []byte{}, len(s)
	for i, j, cur := 0, 0, 0; i < n; i = j {
		j++
		cur++
		for j < n && cur > 0 {
			if s[j] == ')' {
				cur--
			} else {
				cur++
			}
			ans = append(ans, s[j])
			j++
		}
		ans = ans[:len(ans)-1]
	}
	return string(ans)
}
