package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := &Counter{}
	for i := 0; i < 10; i++ {
		go func() {
			for {

				count := c.Read()
				fmt.Println(count)
				time.Sleep(time.Second)
			}

		}()
	}
	for {
		c.Write()
		time.Sleep(time.Second)
	}

}

type Counter struct {
	mu    sync.RWMutex
	Count uint
}

func (c Counter) Read() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Count
}

func (c *Counter) Write() {
	c.mu.Lock()
	c.Count++
	c.mu.Unlock()
}
