package main

import "fmt"

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	isConnected = [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	isConnected = [][]int{{1, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 1, 1}, {0, 0, 1, 1}}
	fmt.Println(findCircleNum(isConnected))
}

func findCircleNumLC(isConnected [][]int) (ans int) {
	vis := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return
}

func findCircleNum(isConnected [][]int) int {
	result := 0
	visited := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(row int) {
		if !visited[row] {
			visited[row] = true
			for idx, conn := range isConnected[row] {
				if conn == 1 && !visited[idx] {
					// 这里需要递归的原因是存在间接相连的可能
					// 例如：[1,0,1,0],[0,1,0,0],[1,0,1,1],[0,0,1,1]
					// 令4个城市为A、B、C、D
					// A -> C; B; C -> D; D -> C;
					// 此时A于D通过C间接相连，省份数量应该是2
					dfs(idx)
				}
			}
		}
	}
	for i, v := range visited {
		if !v {
			result++
			dfs(i)
		}
	}
	return result
}
