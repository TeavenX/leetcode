package main

func main() {

}

func countSubgraphsForEachDiameterCV(n int, edges [][]int) []int {
	// 建树
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1 // 编号改为从 0 开始
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// 求树的直径
	var inSet, vis [15]bool
	var diameter int
	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] && inSet[y] {
				ml := dfs(y) + 1
				diameter = max(diameter, maxLen+ml)
				maxLen = max(maxLen, ml)
			}
		}
		return
	}

	ans := make([]int, n-1)
	var f func(int)
	f = func(i int) {
		if i == n {
			for v, b := range inSet {
				// 随机取链上的一个点，求最大直径
				if b {
					vis, diameter = [15]bool{}, 0
					dfs(v)
					// 因为inSet的每一个点都在链上，所以求完随机一个结果都是一样的，直接break
					break
				}
			}
			// 这里判断相等的目的是？（保证是直径？）
			if diameter > 0 && vis == inSet {
				ans[diameter-1]++
			}
			return
		}

		// 不选城市 i
		f(i + 1)

		// 选城市 i
		inSet[i] = true
		f(i + 1)
		inSet[i] = false // 恢复现场
	}
	f(0)
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 理解了30%
func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	// 建图
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var inSet, vis [15]bool
	var diameter int
	var dfs func(x int) int
	dfs = func(x int) int {
		maxLen := 0
		vis[x] = true
		for _, y := range g[x] {
			if !vis[y] && inSet[y] {
				l := dfs(y) + 1
				diameter = max(diameter, maxLen+l)
				maxLen = max(maxLen, l)
			}
		}
		return maxLen
	}

	ans := make([]int, n-1)
	var search func(i int)
	search = func(i int) {
		if i == n {
			for x, exist := range inSet {
				if exist {
					vis, diameter = [15]bool{}, 0
					dfs(x)
					break
				}
			}
			if diameter > 0 && vis == inSet {
				ans[diameter-1]++
			}
			return
		}
		search(i + 1)
		inSet[i] = true
		search(i + 1)
		inSet[i] = false
	}
	search(0)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
