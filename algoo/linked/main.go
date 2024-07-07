package main

import "fmt"

func main() {
	n1 := NewListNode(1)
	n2 := NewListNode(2)
	n3 := NewListNode(3)
	p := NewListNode(999)
	n1.Next = n2
	n2.Next = n3

	insertNode(n1, p)
	removeItem(n1)
	index := access(n1, 2)
	fmt.Println(index.Val, "visitValue")
	fmt.Println(findNode(n1, 1), "findindex")
}

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 为了修改表结构
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}
func insertNode(n0 *ListNode, p *ListNode) {
	temp := n0.Next
	p.Next = temp
	n0.Next = p
}

func removeItem(n0 *ListNode) {
	if n0.Next == nil {
		return
	}
	d := n0.Next
	n1 := d.Next
	n0.Next = n1
}

// 访问
func access(f *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if f == nil {
			return nil
		}
		f = f.Next

	}
	return f
}

func findNode(f *ListNode, target int) int {
	i := 0
	for f != nil {
		if f.Val == target {
			return i
		}
		f = f.Next
		i++
	}
	return -1
}
