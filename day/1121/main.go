package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	mu    sync.RWMutex
	age   int
	level int
}

func (p *Person) read(i int) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	fmt.Printf("read age: %d,level: %d\n", p.age, p.level)
	if p.level > 0 {
		fmt.Printf("%d 号英雄升级成功", i)
		fmt.Printf("read age: %d,level: %d\n", p.age, p.level)
	}

}
func (p *Person) write() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.age++
	p.level += 10
	fmt.Printf("write age: %d,level: %d\n", p.age, p.level)

}
func main() {
	p := &Person{}
	go p.read(1)
	go p.read(2)
	go p.read(3)
	go p.write()
	// runtime.Gosched()
	time.Sleep(time.Second)
}
