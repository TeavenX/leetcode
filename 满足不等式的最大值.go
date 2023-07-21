package main

import "math"

func main() {

}

type Pair struct {
	x  int
	yx int
}

func findMaxValueOfEquation(points [][]int, k int) int {
	ans := math.MinInt
	pq := make([]Pair, 0)
	for _, p := range points {
		x, y := p[0], p[1]
		for len(pq) > 0 && pq[0].x < x-k {
			pq = pq[1:]
		}
		if len(pq) > 0 {
			ans = max(ans, x+y+pq[0].yx)
		}
		for len(pq) > 0 && pq[len(pq)-1].yx <= y-x {
			pq = pq[:len(pq)-1]
		}
		pq = append(pq, Pair{x, y - x})
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
