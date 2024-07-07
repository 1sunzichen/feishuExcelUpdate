package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello(i int) {

	fmt.Println("hello", i)
}
func hellowg(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hellowg", i)
}

var wg sync.WaitGroup

func main() {

	test1()
}
func test2() {

	for i := 0; i < 2; i++ {

		go hello(i)
		go func(b int) {
			fmt.Println("func", b)
		}(i)
		wg.Add(1)
		go hellowg(i, &wg)
	}
	// time.Sleep(time.Second)
	wg.Wait()
}
func test1() {
	//cpu 核数 Translate a line under cursor
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()

}
func a() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("A%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("B%d\n", i)
	}
}
