package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(node int) *TreeNode {
	return &TreeNode{
		Val:   node,
		Left:  nil,
		Right: nil,
	}
}

var numdeep []any

func main() {
	test()

	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n1.Left = n2
	n2.Left = n3

	// frontDeep(n1)
	var numT1 []any
	frontDeepT1(n1, &numT1)
	fmt.Println(numT1...)
	// fmt.Println(lenvelOrder(n1)...)
}

// 广度优先遍历
// 声明 list队列 。声明切片，放入根节点。
// for循环条件 如果 队列长度大于0  ， 将前面的值存入切片，判断 左节点和右节点，有值放入队列等待遍历⌛️。

func lenvelOrder(r *TreeNode) []any {
	quene := list.New()
	quene.PushBack(r)
	nums := make([]any, 0)
	for quene.Len() > 0 {
		node := quene.Remove(quene.Front()).(*TreeNode)
		nums = append(nums, node.Val)
		if node.Left != nil {
			quene.PushBack(node.Left)
		}
		if node.Right != nil {
			quene.PushBack(node.Right)
		}

	}
	return nums
}

func frontDeep(t *TreeNode) {
	if t == nil {
		return
	}
	// 前序遍历  根 左 右
	// 后序遍历  左 右 根
	// 中序遍历  左 根 右
	numdeep = append(numdeep, t.Val) //根

	frontDeep(t.Left) //左

	frontDeep(t.Right) //右
}
func test() {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)
}
func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
		i++
	}
	*/

	for i := range s {
		//这里修改array 动了指针 就会影响到原来的 切片
		s[i] += 1
	}
}

// 传入指针即可    append 复制 不会修改外层的值。传指针切片的时候可以修改。 append 相当于 复制拷贝 值传递
// 不用append 方法 直接修改 切片的值 也可以修改 原来的值。
func frontDeepT1(t *TreeNode, numd *[]any) {
	if t == nil {
		return
	}
	// 前序遍历  根 左 右
	// 后序遍历  左 右 根
	// 中序遍历  左 根 右
	*numd = append(*numd, t.Val) //根

	frontDeepT1(t.Left, numd) //左
	if t.Right != nil {

		frontDeepT1(t.Right, numd) //右
	}
}
