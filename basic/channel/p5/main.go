package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()

	a := recur1(5)
	fmt.Println(a)
	func1()
	// m := make(map[int]bool)
	// for {
	// 	m[rand.Intn(100000000)] = true
	// }

}

func func1() {
	func2()
}

func func2() {
	func3()
}

func func3() {
	m := make(map[int]bool)
	for {
		m[rand.Intn(100000000)] = true
	}
}
func recur(n int) int {
	if n == 1 {
		return 1
	}
	res := recur(n - 1)
	return res + n
}
func recur1(n int) int {
	if n == 1 {
		return 1
	}
	return recur1(n-1) + n
}
