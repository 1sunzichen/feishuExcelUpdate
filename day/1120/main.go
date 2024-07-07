package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Person struct {
	mu    int32
	age   int
	level int
}

func (p *Person) Add() {
	//原子加锁
	atomic.CompareAndSwapInt32(&p.mu, 0, 1)
	p.age += 1
	p.level += 1
	fmt.Printf("age:%d,level:%d\n", p.age, p.level)
	atomic.CompareAndSwapInt32(&p.mu, 1, 0)
}
func main() {
	p := Person{age: 1, level: 1}
	go p.Add()
	go p.Add()
	go p.Add()
	time.Sleep(time.Second)
}
