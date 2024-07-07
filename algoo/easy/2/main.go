// N 叉树的后序遍历
package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func dfsFunc(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, v := range node.Children {
			dfs(v)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}
func main() {
	root := &Node{Val: 2}
	root.Children = []*Node{
		{Val: 2, Children: []*Node{
			{Val: 42},
			{Val: 44},
		}},
		{Val: 3},
		{Val: 4},
	}
	fmt.Println(dfsFunc(root))
}
