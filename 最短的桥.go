package main

func main() {

}

var steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func shortestBridge(grid [][]int) int {
	// 用两个变量标记岛屿的数字，单数表示岛屿1，双数表示岛屿2
	// 从编号为 0 的点出发，向四周延伸，碰到单数/双数的点后，记录状态1，又碰到双数/单数点后，判断路程
	var minLen int = 1e4
	n := len(grid)
	var markIsland func(i, j int)
	markIsland = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 1 {
			grid[i][j] = 2
		} else {
			return
		}
		for _, step := range steps {
			markIsland(i+step[0], j+step[1])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				markIsland(i, j)
				goto next
			}
		}
	}
next:
	search := func(i, j int) {
		dist := [3]int{0, n, n}
		t := i - 1
		for t >= 0 {
			if grid[t][j] > 0 && i-t < dist[grid[t][j]] {
				dist[grid[t][j]] = i - t
				break
			}
			t--
		}
		t = i + 1
		for t < n {
			if grid[t][j] > 0 && t-i < dist[grid[t][j]] {
				dist[grid[t][j]] = t - i
				break
			}
			t++
		}
		t = j - 1
		for t >= 0 {
			if grid[i][t] > 0 && j-t < dist[grid[i][t]] {
				dist[grid[i][t]] = j - t
				break
			}
			t--
		}
		t = j + 1
		for t < n {
			if grid[i][t] > 0 && t-j < dist[grid[i][t]] {
				dist[grid[i][t]] = t - j
				break
			}
			t++
		}
		if dist[1] < n && dist[2] < n {
			minLen = min(dist[1]+dist[2]-1, minLen)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				search(i, j)
			}
		}
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
