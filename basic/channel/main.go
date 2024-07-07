package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// channel 管道是链接每个协程的类型数据
// 当下协程 管道发送数据 一般就会在 其他 协程 管道中接收数据
// 发送数据：主进程 没有缓存 一定要有其他协程 接受数据之后才能发送数据或者单开线程发送数据/有缓存 可以先发送数据到管道的缓存中
var wg sync.WaitGroup

func main() {
	// test()
	// trace.Start(os.Stderr)
	// defer trace.Stop()
	test()
	// test9141()
	// main1001()
}

// 条件：保持主进程未关闭的情况下，开启其他协程（goruntine） 用通道（channel） 进行通信
// 多运行几次 给通道发送值。如果时间片内，没有接受到值。不输出a 运行主协程
func test() {

	ch1 := make(chan int)
	go func() {
		a := <-ch1
		//阻塞了
		fmt.Println("goruntine1:", a)
	}()
	//给通道发送值。如果时间片内，没有接受到值。不输出a 运行主协程
	ch1 <- 9
	// 获取当前 goroutine 数量
	numGoroutines := runtime.NumGoroutine()
	fmt.Println("Number of goroutines:", numGoroutines)
	fmt.Println("maingoruntine:send", ch1)
}

// 运行主协程运行完 还没到在主进程时间片执行完所有程序，就没法 输出a值
// 思考？：进程 线程越多，切换成本大 。基本无法输出a
func test00() {

	ch1 := make(chan int)

	go func() {
		a := <-ch1

		fmt.Println("goruntine1:", a)

	}()
	go func() {
		ch1 <- 9
		fmt.Println("goruntine2:send", ch1)
	}()
	fmt.Println("maingoruntine:send", ch1)
}

// 运行主协程运行完 还没到
func test000() {

	ch1 := make(chan int)

	go func() {
		a := <-ch1

		fmt.Println("goruntine1:", a)

	}()
	go func() {
		ch1 <- 9
		fmt.Println("goruntine2:send", ch1)
	}()
	//让出时间片
	runtime.Gosched()
	fmt.Println("maingoruntine:send", ch1)

}
func test11() {
	wg.Add(1)
	ch1 := make(chan []int)
	go func() { ch1 <- []int{9} }()
	go func() {
		defer wg.Done()
		a := <-ch1
		// fmt.Printf("%T,%d", a, a[0])
		fmt.Println(a[0])
	}()

	wg.Wait()
	close(ch1)
}

func test1() {
	//
	wg.Add(1)
	ch1 := make(chan []int)
	go func() {
		defer wg.Done()
		a := <-ch1
		// fmt.Printf("%T,%d", a, a[0])
		fmt.Println(a[0])
	}()

	ch1 <- []int{9}
	wg.Wait()
	close(ch1)
}
func test2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {

			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i
		}
		close(ch2)
	}()
	for v := range ch2 {
		fmt.Println(v)
	}
}
func test9141() {
	wg.Add(2)
	a := make(chan int, 100)
	b := make(chan int, 100)

	go test914channel(a)
	go test914channelreve(a, b)
	wg.Wait()
	for v := range b {
		fmt.Println(v)
	}

}
func test914channel(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
		fmt.Println("ich1", i)
	}
	close(ch1)
}
func test914channelreve(ch1, ch2 chan int) {
	defer wg.Done()
	for v := range ch1 {

		ch2 <- v

	}
	close(ch2)
}

var dog = make(chan struct{})
var cat = make(chan struct{})
var fish = make(chan struct{})

func Dog() {
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
}

func Cat() {
	<-dog
	fmt.Println("cat")
	cat <- struct{}{}
}

func Fish() {
	<-cat
	fmt.Println("fish")
	fish <- struct{}{}
}

func main1001() {
	for i := 0; i < 100; i++ {
		go Cat()
		go Fish()
		go Dog()
	}
	//决定顺序
	dog <- struct{}{}

	time.Sleep(10 * time.Second)
}
