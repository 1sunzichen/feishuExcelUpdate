package main

import "fmt"

func If(c bool, t, f interface{}) interface{} {
	if c {
		return t
	}
	return f
}
func main() {
	a, b := 2, 13
	max := If(a > b, a, b).(int)
	fmt.Println(max, "max")
}
