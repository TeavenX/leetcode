package main

import "fmt"

func main() {
	dictionary := []string{"leet", "code", "leetcode"}
	fmt.Println(minExtraChar("leetscode", dictionary))
}

func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	cache := make([]int, n)
	var dfs func(idx int) int
	dfs = func(idx int) (res int) {
		if idx == n {
			return 0
		}
		if v := cache[idx]; v > 0 {
			return v
		}
		defer func() {
			cache[idx] = res
		}()
		ans := 0
		for i := idx; i < n; i++ {
			res := 0
			for _, dict := range dictionary {
				l := len(dict)
				if i+l > n {
					continue
				}
				if s[i:i+l] == dict {
					res = max(res, dfs(i+l)+l)
				}
			}
			ans = max(ans, res)
		}
		return ans
	}
	return n - dfs(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
