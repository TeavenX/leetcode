package main

func main() {

}

func maximalNetworkRank(n int, roads [][]int) int {
	// 建立邻接矩阵
	deg := make([][]int, n)
	// 统计每个节点的度（由于是无向图，所以等同于出现次数）
	count := make([]int, n)
	ans := 0
	for i := range deg {
		deg[i] = make([]int, n)
	}
	for _, r := range roads {
		// 这里设为1方便后面计算秩
		deg[r[0]][r[1]] = 1
		deg[r[1]][r[0]] = 1
		count[r[0]]++
		count[r[1]]++
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// 由于两个直接相连的节点的秩只能算1次，所以如果两个点相连这里需要减去1
			ans = max(ans, count[i]+count[j]-deg[i][j])
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
