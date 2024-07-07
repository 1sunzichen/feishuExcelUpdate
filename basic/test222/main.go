package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	slice1 := []int{1, 2, 3}
	s2 := append(slice1, 2)
	fmt.Println(cap(slice1), cap(s2))
	h := convertToLinkedList(slice1)
	ForH(h)
}
func ForH(list *ListNode) {
	fmt.Println(list.Val)
	if list.Next != nil {
		ForH(list.Next)
	}

}

// slice 转链表
func convertToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}

		current.Next = node
		//完成置换 更新最新的节点
		current = node
	}
	return head
}
