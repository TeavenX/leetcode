package main

func main() {

}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			t := min(rowSum[i], colSum[j])
			ans[i][j] = t
			rowSum[i] -= t
			colSum[j] -= t
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
