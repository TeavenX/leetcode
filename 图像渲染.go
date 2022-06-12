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
