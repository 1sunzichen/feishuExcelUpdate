package main

import "fmt"

func lengthOfLastWord(s string) (ans int) {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		index--
		ans++
	}
	return
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
}
