package main

func main() {

}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	tree := make([][]int, n)
	for i, m := range manager {
		if m >= 0 {
			tree[m] = append(tree[m], i)
		}
	}
	var dfs func(head int) int
	dfs = func(head int) int {
		if len(tree[head]) == 0 {
			return 0
		}
		mx := informTime[head]
		t := 0
		for _, i := range tree[head] {
			t = max(t, dfs(i))
		}
		return mx + t
	}
	return dfs(headID)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func numOfMinutesV2(n int, headID int, manager []int, informTime []int) int {
	cache := make([]int, n)
	for _, i := range manager {
		if i >= 0 {
			cache[i] = -1
		}
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if manager[i] < 0 {
			return informTime[i]
		}
		if cache[i] >= 0 {
			return cache[i]
		}
		res := dfs(manager[i]) + informTime[i]
		cache[i] = res
		return res
	}
	ans := 0
	for i := range manager {
		ans = max(ans, dfs(i))
	}
	return ans
}
