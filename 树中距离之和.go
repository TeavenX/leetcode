package main

func main() {

}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	ans := make([]int, n)
	g := make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	size := make([]int, n)
	var dfs func(p int, cur int, depth int)
	dfs = func(p, cur, depth int) {
		ans[0] += depth
		size[cur] = 1
		for _, nxt := range g[cur] {
			if nxt != p {
				dfs(cur, nxt, depth+1)
				size[cur] += size[nxt]
			}
		}
	}
	dfs(-1, 0, 0)
	var reRoot func(p int, cur int)
	reRoot = func(p, cur int) {
		ans[cur] = ans[p] + n - size[cur] - size[cur]
		for _, nxt := range g[cur] {
			if nxt != p {
				reRoot(cur, nxt)
			}
		}
	}
	for _, cur := range g[0] {
		reRoot(0, cur)
	}
	return ans
}
