package main

import "sort"

func main() {

}

type Pair struct {
	x, y int
}

func countPairs(n int, edges [][]int, queries []int) []int {
	d := make([]int, n+1)
	count := make(map[Pair]int)
	for _, e := range edges {
		x, y := e[0], e[1]
		if x > y {
			x, y = y, x
		}
		d[x]++
		d[y]++
		count[Pair{x, y}]++
	}
	sorted := append([]int{}, d...)
	sort.Ints(sorted)
	ans := make([]int, len(queries))
	for i, q := range queries {
		left, right := 1, n
		for left < right {
			if sorted[left]+sorted[right] <= q {
				left++
			} else {
				ans[i] += right - left
				right--
			}
		}
		for e, c := range count {
			if d[e.x]+d[e.y] > q && d[e.x]+d[e.y]-c <= q {
				ans[i]--
			}
		}
	}
	return ans
}
