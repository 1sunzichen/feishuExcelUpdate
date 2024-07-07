package main

import (
	"container/list"
	"fmt"
)

// 基于链表的实现
type LinkedListstack struct {
	data *list.List
}

func newLinkedListStack() *LinkedListstack {
	return &LinkedListstack{
		data: list.New(),
	}
}
func (s *LinkedListstack) push(value int) {
	s.data.PushBack(value)
}
func (s *LinkedListstack) pop() any {
	if s.data.Len() == 0 {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}
func testLinkedListstack() {
	a := newLinkedListStack()
	a.push(1)
	a.push(2)
	b := a.pop()
	fmt.Println(b)
}
func main() {
	// testLinkedListstack()
}

// 基于数组的实现
type arraystack struct {
	data []int
}

func newArrayStck() *arraystack {
	return &arraystack{
		data: make([]int, 0, 16),
	}
}
