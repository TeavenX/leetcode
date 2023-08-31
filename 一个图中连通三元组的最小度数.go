package main

import "math"

func main() {

}

func minTrioDegree(n int, edges [][]int) int {
	g := make([][]bool, n+1)
	for i := range g {
		g[i] = make([]bool, n+1)
	}
	deg := make([]int, n+1)
	for _, e := range edges {
		g[e[0]][e[1]] = true
		g[e[1]][e[0]] = true
		deg[e[0]]++
		deg[e[1]]++
	}
	ans := math.MaxInt
	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			if g[a][b] {
				for c := b + 1; c <= n; c++ {
					if g[a][c] && g[b][c] {
						ans = min(ans, deg[a]+deg[b]+deg[c]-6)
					}
				}
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
