package main

import (
	"math"
)

func main() {
	// [
	// 0  [0,0,1,0,0,0,0,0,0,0,0,0,0,0,0],
	// 1  [0,1,0,1,1,0,0,1,0,0,0,0,1,0,0],
	// 2  [0,1,0,0,0,0,1,0,0,1,0,0,0,0,0],
	// 3  [0,0,0,0,0,0,1,0,0,0,0,0,0,0,0],
	// 4  [0,0,0,0,0,0,1,1,0,0,0,0,0,0,0],
	// 5  [0,0,0,0,0,0,0,0,0,1,0,1,0,0,0],
	// 6  [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
	// 7  [0,0,0,1,0,1,0,0,1,0,0,0,1,0,0],
	// 8  [0,0,0,0,1,0,0,0,0,0,0,0,0,1,0],
	// 9  [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
	// 10 [0,0,0,1,0,0,0,0,0,0,0,0,0,0,0],
	// 11 [1,0,1,1,0,0,0,0,0,0,0,0,0,0,0],
	// 12 [0,0,0,0,0,0,0,0,0,0,0,0,0,1,0],
	// 13 [1,0,0,0,0,0,1,0,0,0,1,0,0,0,1],
	// 14 [0,0,1,0,1,0,0,0,0,0,0,0,0,0,0]
	// ]

}

type point struct {
	x1, y1, x2, y2 int
	step           int
}

type cn struct {
	x1, y1, x2, y2 int
}

func minimumMoves(grid [][]int) int {
	n := len(grid)
	ans := math.MaxInt
	cache := map[cn]bool{{0, 0, 0, 1}: true}
	q := []point{{0, 0, 0, 1, 0}}
	valid := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < n && y < n && grid[x][y] != 1
	}
	walk := func(x1, y1, x2, y2, step int) {
		if valid(x1, y1) && valid(x2, y2) {
			if x1 == n-1 && y1 == n-2 && x2 == n-1 && y2 == n-1 {
				ans = min(ans, step+1)
			}
			p := cn{x1, y1, x2, y2}
			if !cache[p] {
				q = append(q, point{x1, y1, x2, y2, step + 1})
				cache[p] = true
			}
		}
	}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			snk := q[i]
			walk(snk.x1+1, snk.y1, snk.x2+1, snk.y2, snk.step)
			walk(snk.x1, snk.y1+1, snk.x2, snk.y2+1, snk.step)
			if snk.x1 == snk.x2 && snk.x1 < n-1 && grid[snk.x1+1][snk.y1+1] != 1 {
				walk(snk.x1, snk.y1, snk.x1+1, snk.y1, snk.step)
			}
			if snk.y1 == snk.y2 && snk.y1 < n-1 && grid[snk.x1+1][snk.y1+1] != 1 {
				walk(snk.x1, snk.y1, snk.x1, snk.y1+1, snk.step)
			}
		}
		q = q[size:]
	}
	if ans == math.MaxInt {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
