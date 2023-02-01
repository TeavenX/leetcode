package main

func main() {

}

type pair struct {
	idx, dist, color int
}

type key struct {
	idx, color int
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	ans := make([]int, n)
	cr, cb := make([][]bool, n), make([][]bool, n)
	visit := make(map[key]bool)
	for i := 0; i < n; i++ {
		cr[i] = make([]bool, n)
		cb[i] = make([]bool, n)
		ans[i] = 400 * n
	}
	for _, e := range redEdges {
		cr[e[0]][e[1]] = true
	}
	for _, e := range blueEdges {
		cb[e[0]][e[1]] = true
	}
	ans[0] = 0
	q := []pair{{0, 0, 1}, {0, 0, 0}}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			pre := q[i]
			for j := 0; j < n; j++ {
				v := key{j, pre.color ^ 1}
				if !visit[v] && (cb[pre.idx][j] && v.color != 0 || cr[pre.idx][j] && v.color != 1) {
					visit[v] = true
					p := pair{j, pre.dist + 1, pre.color ^ 1}
					ans[j] = min(ans[j], pre.dist+1)
					q = append(q, p)
				}
			}
		}
		q = q[size:]
	}
	for i, num := range ans {
		if num == n*400 {
			ans[i] = -1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
