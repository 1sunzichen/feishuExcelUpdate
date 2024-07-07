package main

import "fmt"

type Graph struct {
	id        int
	neighbors []*Graph
}

func (g *Graph) DFS(visited map[*Graph]bool) {
	if visited[g] {
		return
	}
	visited[g] = true
	fmt.Println(g.id)
	for _, neighbor := range g.neighbors {
		neighbor.DFS(visited)
	}
}
func app(b *[]int, c int) {

	*b = append(*b, c)
}
func app2(b []int, c int) {
	fmt.Println(b)
	b = append(b, c)
}
func main() {
	a := []int{1, 2, 3}
	fmt.Println(a, "a")
	// app(&a, 3)
	a = append(a, 3)
	// app2(a, 3)

	fmt.Println(a, "a")
	node1 := &Graph{id: 1}
	node2 := &Graph{id: 2}
	node3 := &Graph{id: 3}
	node4 := &Graph{id: 4}

	node1.neighbors = []*Graph{node2, node3}
	node2.neighbors = []*Graph{node1, node4}
	node3.neighbors = []*Graph{node1}
	node4.neighbors = []*Graph{node2}

	// 创建一个空的visited映射
	visited := make(map[*Graph]bool)
	node1.DFS(visited)

}
