package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Time
	pow := time.Now()
	fn1(pow)

}
func fn3() {
	pow := time.Now()
	fmt.Printf("%#v\n", pow)
	time1 := pow.Unix()
	fmt.Printf("%v\n", time1)
	// tickDemo()
	fmt.Println(pow.Format("2006-01-02 15:04:05"))
	nextYear, err := time.Parse("2006-01-02", "2020-08-03")
	if err != nil {
		return
	}
	d := pow.Sub(nextYear)
	fmt.Println(d)
}

func tickDemo() {
	//定时器
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Print(i)
	}
}

func fn1(now time.Time) {
	str := "2006-01-02 15:04:05"
	str1 := "2023-09-01 14:41:02"
	// time.Parse(str, str1)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load  loc failed,err:%v\n", err)
		return
	}
	timeObj, err := time.ParseInLocation(str, str1, loc)
	if err != nil {
		fmt.Printf("load  loc failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	td := timeObj.Sub(now)
	fmt.Println(td)

}
