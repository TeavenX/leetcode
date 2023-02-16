package main

func main() {

}

func largest1BorderedSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	rs := make([][]int, n)
	cs := make([][]int, m)
	for i := range rs {
		rs[i] = make([]int, m+1)
	}
	for i := range cs {
		cs[i] = make([]int, n+1)
	}
	for i, row := range grid {
		for j, num := range row {
			rs[i][j+1] = rs[i][j] + num
			cs[j][i+1] = cs[j][i] + num
		}
	}
	for d := min(n, m); d > 0; d-- {
		for i := 0; i <= n-d; i++ {
			for j := 0; j <= m-d; j++ {
				if rs[i][j+d]-rs[i][j] == d &&
					cs[j][i+d]-cs[j][i] == d &&
					rs[i+d-1][j+d]-rs[i+d-1][j] == d &&
					cs[j+d-1][i+d]-cs[j+d-1][i] == d {
					return d * d
				}
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
