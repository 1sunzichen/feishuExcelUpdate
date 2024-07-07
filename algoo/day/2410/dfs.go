package main

import "fmt"

func dfs(grid [][]int, i, j int, visited [][]bool) {
	var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dirCore func(i, j int)
	dirCore = func(i, j int) {
		m, n := len(grid), len(grid[0])
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		for _, d := range dirs {
			dirCore(i+d[0], j+d[1])
		}
	}
	dirCore(i, j)
}
func main() {
	var ddd = [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	v := make([][]bool, 5)
	for i := range v {
		v[i] = make([]bool, 5)
	}

	dfs(ddd, 0, 0, v)
	fmt.Println(ddd, v)
}
