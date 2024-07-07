package main

import "fmt"

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(Val int) *TreeNode {
	return &TreeNode{
		Val:   Val,
		Left:  nil,
		Right: nil,
	}
}
func preOrder(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	if root.Val.(int) == 7 {
		*res = append(*res, root)
	}
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

func preOrderII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	if root == nil {
		return
	}
	*path = append(*path, root)
	if root.Val.(int) == 7 {
		*res = append(*res, *path)
	}
	preOrderII(root.Left, res, path)
	preOrderII(root.Right, res, path)
	*path = (*path)[:len(*path)-1]
}
func main() {
	node := NewTreeNode(1)
	node.Left = &TreeNode{Val: 2}
	node.Right = &TreeNode{Val: 3}
	node.Left.Left = &TreeNode{Val: 4}
	node.Left.Right = &TreeNode{Val: 7}
	node.Right.Right = &TreeNode{Val: 7}
	res := new([]*TreeNode)
	preOrder(node, res)
	for k, v := range *res {
		fmt.Println(v.Val, k)
	}
	res2 := new([][]*TreeNode)
	path := new([]*TreeNode)
	preOrderII(node, res2, path)
	for k, v := range *path {
		fmt.Println(v.Val, k)
	}
	fmt.Println("\n回溯")
	for _, v := range *res2 {
		fmt.Print("[")
		for kk, vv := range v {
			if kk != len(v)-1 {
				fmt.Print(vv.Val, ",")
			} else {
				fmt.Print(vv.Val)
			}
		}
		fmt.Println("]")
	}

	fmt.Println("\n剪枝")

}
