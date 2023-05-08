package main

func main() {

}

var steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func minPushBox(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n*m)
	for i := range visited {
		visited[i] = make([]bool, n*m)
	}
	var si, sj, bi, bj int
	for i, row := range grid {
		for j, b := range row {
			if b == 'S' {
				si, sj = i, j
			}
			if b == 'B' {
				bi, bj = i, j
			}
		}
	}
	calc := func(x, y int) int {
		return x*m + y
	}
	check := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m && grid[x][y] != '#'
	}
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{calc(si, sj), calc(bi, bj), 0})
	visited[calc(si, sj)][calc(bi, bj)] = true
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		si, sj, bi, bj := pos[0]/m, pos[0]%m, pos[1]/m, pos[1]%m
		d := pos[2]
		if grid[bi][bj] == 'T' {
			return d
		}
		for _, step := range steps {
			sx, sy := si+step[0], sj+step[1]
			if !check(sx, sy) {
				continue
			}
			if sx == bi && sy == bj {
				bx, by := bi+step[0], bj+step[1]
				if check(bx, by) && !visited[calc(sx, sy)][calc(bx, by)] {
					visited[calc(sx, sy)][calc(bx, by)] = true
					queue = append(queue, [3]int{calc(sx, sy), calc(bx, by), d + 1})
				}
			} else if !visited[calc(sx, sy)][calc(bi, bj)] {
				visited[calc(sx, sy)][calc(bi, bj)] = true
				queue = append([][3]int{{calc(sx, sy), calc(bi, bj), d}}, queue...)
			}
		}
	}
	return -1
}
