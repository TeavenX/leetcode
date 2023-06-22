package main

import "sort"

func main() {

}

var steps = [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func pondSizes(land [][]int) []int {
	n, m := len(land), len(land[0])
	valid := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m && land[i][j] == 0
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if !valid(i, j) {
			return 0
		}
		land[i][j] = 1000
		ans := 1
		for _, s := range steps {
			ans += dfs(i+s[0], j+s[1])
		}
		return ans
	}
	ans := make([]int, 0)
	for i, row := range land {
		for j, v := range row {
			if v == 0 {
				ans = append(ans, dfs(i, j))
			}
		}
	}
	sort.Ints(ans)
	return ans
}
