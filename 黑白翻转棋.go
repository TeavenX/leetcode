package main

func main() {

}

var steps = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func flipChess(chessboard []string) int {
	n, m := len(chessboard), len(chessboard[0])
	valid := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m
	}
	bfs := func(i, j int) int {
		g := make([][]byte, n)
		for i := range g {
			g[i] = []byte(chessboard[i])
		}
		g[i][j] = 'X'
		q := []int{i*10 + j}
		count := 0
		for len(q) > 0 {
			i, j = q[0]/10, q[0]%10
			q = q[1:]
			for _, s := range steps {
				a, b := s[0], s[1]
				x, y := i+a, j+b
				for valid(x, y) && g[x][y] == 'O' {
					x += a
					y += b
				}
				if valid(x, y) && g[x][y] == 'X' {
					x -= a
					y -= b
					count += max(abs(x-i), abs(y-j))
					for x != i || y != j {
						g[x][y] = 'X'
						q = append(q, x*10+y)
						x -= a
						y -= b
					}
				}
			}
		}
		return count
	}
	ans := 0
	for i, r := range chessboard {
		for j, v := range r {
			if v == '.' {
				ans = max(ans, bfs(i, j))
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
