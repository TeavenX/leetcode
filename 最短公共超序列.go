package main

func main() {

}

func shortestCommonSupersequenceOOM(str1 string, str2 string) (ans string) {
	n, m := len(str1), len(str2)

	cache := make([][]string, n)
	for i := range cache {
		cache[i] = make([]string, m)
	}

	var dfs func(i, j int) string
	dfs = func(i, j int) (ans string) {
		if i < 0 {
			return str2[:j+1]
		}
		if j < 0 {
			return str1[:i+1]
		}
		if v := cache[i][j]; v != "" {
			return v
		}
		defer func() {
			cache[i][j] = ans
		}()
		if str1[i] == str2[j] {
			return dfs(i-1, j-1) + string(str1[i])
		}
		a1 := dfs(i-1, j) + string(str1[i])
		a2 := dfs(i, j-1) + string(str2[j])
		if len(a1) < len(a2) {
			return a1
		}
		return a2
	}
	return dfs(n-1, m-1)
}

func shortestCommonSupersequence(str1 string, str2 string) (ans string) {
	n, m := len(str1), len(str2)

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, m)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) (ans int) {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if v := cache[i][j]; v != 0 {
			return v
		}
		defer func() {
			cache[i][j] = ans
		}()
		if str1[i] == str2[j] {
			return dfs(i-1, j-1) + 1
		}
		return min(dfs(i-1, j), dfs(i, j-1)) + 1
	}

	var makeStr func(i, j int) string
	makeStr = func(i, j int) string {
		if i < 0 {
			return str2[:j+1]
		}
		if j < 0 {
			return str1[:i+1]
		}
		if str1[i] == str2[j] {
			return makeStr(i-1, j-1) + string(str1[i])
		}
		if dfs(i, j) == dfs(i-1, j)+1 {
			return makeStr(i-1, j) + string(str1[i])
		}
		return makeStr(i, j-1) + string(str2[j])
	}
	return makeStr(n-1, m-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
