package main

import "fmt"

// 输出什么
func SelectExaml() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	c1 <- 1
	c2 <- 2
	select {
	case <-c1:
		fmt.Println("c1")
	case <-c2:
		fmt.Println("c2")
	}
}

// 输出
func SelectExam2() {
	c := make(chan int)
	select {
	case <-c:
		fmt.Printf("read")
	case c <- 1:
		fmt.Println("writ")
	}
}

// 关闭的通道 可以读取
func SelectExam3() {
	c := make(chan int, 10)
	c <- 1
	close(c)
	select {
	case d := <-c:
		fmt.Println(d)
	}
}

func SelectExam4() {

	c := make(chan int, 10)
	c <- 1
	close(c)
	select {
	case d, ok := <-c:
		if !ok {
			fmt.Println("no data")
			break
		}
		fmt.Println(d)
	}
}

// 阻塞
func SelectExam5() {
	select {}
}

// default
func SelectExam6() {
	var c chan string
	select {
	case c <- "Hello":
		fmt.Println("sent")
	default:
		fmt.Println("default")
	}
}

// 读操作 3 个case都有机会执行
func SelectAssign(c chan string) {
	select {
	case <-c: //0个变量
		fmt.Printf("0")
	case d := <-c:
		fmt.Printf("1 received%s\n", d)
	case d, ok := <-c:
		if !ok {
			fmt.Printf("no data found\n")
			break
		}
		fmt.Printf("2:reve%s\n", d)
	}
}
func main() {
	// SelectExam6()
	c := make(chan string)
	close(c)
	SelectAssign(c)
}

//select
