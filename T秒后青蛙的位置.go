package main

func main() {

}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	g := make([][]int, n+1)
	for i := range g {
		g[i] = make([]int, 0)
	}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	vis := make([]bool, n+1)
	var dfs func(t int, pos int) float64
	dfs = func(t, pos int) float64 {
		if vis[pos] {
			return 0
		}
		if pos == target {
			if t == 0 {
				return 1
			}
			for _, p := range g[pos] {
				if !vis[p] {
					return 0
				}
			}
			return 1
		}
		if t == 0 {
			return 0
		}
		vis[pos] = true
		count, total := 0, 0
		var psb float64
		for _, p := range g[pos] {
			if !vis[p] {
				total++
			}
			if v := dfs(t-1, p); v > 0 {
				count++
				psb += v
			}
		}
		vis[pos] = false
		psb = psb * float64(count) / float64(total)
		return psb
	}
	return dfs(t, 1)
}
