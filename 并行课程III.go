package main

func main() {

}

func minimumTime(n int, relations [][]int, time []int) int {
	g := make([][]int, n+1)
	indeg := make([]int, n+1)
	for _, r := range relations {
		g[r[0]] = append(g[r[0]], r[1])
		indeg[r[1]]++
	}
	q := make([]int, 0)
	f := make([]int, n+1)
	ans := 0
	for i, nd := range indeg {
		if i > 0 && nd == 0 {
			q = append(q, i)
			f[i] = time[i-1]
			ans = max(ans, f[i])
		}
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range g[p] {
			indeg[nxt]--
			if indeg[nxt] == 0 {
				q = append(q, nxt)
			}
			f[nxt] = max(f[nxt], f[p]+time[nxt-1])
			ans = max(ans, f[nxt])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
