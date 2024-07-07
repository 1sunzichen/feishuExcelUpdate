package main1

import "fmt"

func main() {

}
func Add(a, b int) int {
	return a + b
}
func PrintAdd(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func PrintString(a ...string) {
	for _, v := range a {
		fmt.Println(v)
	}
}
