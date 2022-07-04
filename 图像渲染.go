package main

import "fmt"

func main() {
	image := make([][]int, 3)
	for i := 0; i < 3; i++ {
		image[i] = make([]int, 3)
	}
	image[0][0] = 1
	image[0][1] = 1
	image[0][2] = 1
	image[1][0] = 1
	image[1][1] = 1
	image[2][0] = 1
	image[2][2] = 1

	result := floodFill(image, 1, 1, 2)
	for _, i := range result {
		fmt.Println(i)
	}
}

type QueueNode struct {
	x, y int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	queue := make([]QueueNode, 0)
	initColor := image[sr][sc]
	queue = append(queue, QueueNode{sr, sc})
	if initColor == newColor {
		return image
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		image[cur.x][cur.y] = newColor
		if cur.x+1 < len(image) && image[cur.x+1][cur.y] == initColor {
			queue = append(queue, QueueNode{cur.x + 1, cur.y})
		}
		if cur.x-1 >= 0 && image[cur.x-1][cur.y] == initColor {
			queue = append(queue, QueueNode{cur.x - 1, cur.y})
		}
		if cur.y+1 < len(image[0]) && image[cur.x][cur.y+1] == initColor {
			queue = append(queue, QueueNode{cur.x, cur.y + 1})
		}
		if cur.y-1 >= 0 && image[cur.x][cur.y-1] == initColor {
			queue = append(queue, QueueNode{cur.x, cur.y - 1})
		}
	}
	return image
}

var steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var dfs func(sr, sc int)
	preColor := image[sr][sc]
	n, m := len(image), len(image[0])
	visited := make(map[string]bool)
	dfs = func(sr, sc int) {
		if visited[fmt.Sprintf("%d,%d", sr, sc)] {
			return
		}
		if sr < 0 || sc < 0 || sr >= n || sc >= m {
			return
		}
		if image[sr][sc] != preColor {
			return
		}
		image[sr][sc] = color
		visited[fmt.Sprintf("%d,%d", sr, sc)] = true
		for _, step := range steps {
			dfs(sr+step[0], sc+step[1])
		}
	}
	dfs(sr, sc)
	return image
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var dfs func(sr, sc int)
	preColor := image[sr][sc]
	if preColor == color {
		return image
	}
	n, m := len(image), len(image[0])
	dfs = func(sr, sc int) {
		if sr < 0 || sc < 0 || sr >= n || sc >= m {
			return
		}
		if image[sr][sc] != preColor {
			return
		}
		image[sr][sc] = color
		for _, step := range steps {
			dfs(sr+step[0], sc+step[1])
		}
	}
	dfs(sr, sc)
	return image
}
