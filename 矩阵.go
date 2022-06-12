package main

import "fmt"

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	mat = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	result := updateMatrix(mat)
	fmt.Println(result)
}

type node struct {
	x, y int
}

func updateMatrix(mat [][]int) [][]int {
	dirs := []node{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(mat), len(mat[0])
	queue := make([]node, 0)
	visited := make([][]bool, m)
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				visited[i][j] = true
				queue = append(queue, node{i, j})
			}
		}
	}
	for len(queue) > 0 {
		lenN := len(queue)
		for i := 0; i < lenN; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				x, y := cur.x+dir.x, cur.y+dir.y
				if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] {
					visited[x][y] = true
					result[x][y] = result[cur.x][cur.y] + 1
					queue = append(queue, node{x, y})
				}
			}
		}
	}
	return result
}
