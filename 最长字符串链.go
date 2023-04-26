package main

import "sort"

func main() {

}

func longestStrChain(words []string) int {
	n := len(words)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	cache := make(map[string]int)
	var dfs func(i int, pre string) int
	dfs = func(i int, pre string) int {
		if i == n {
			return 0
		}
		if v, ok := cache[pre]; ok {
			return v
		}
		res := 0
		for ; i < n && (len(words[i])-len(pre) <= 1 || pre == ""); i++ {
			cur := words[i]
			if cur == pre {
				continue
			}
			p1, p2 := 0, 0
			for p1 < len(pre) && p2 < len(cur) {
				if cur[p2] != pre[p1] {
					p2++
					continue
				}
				p1++
				p2++
			}
			if p1 == len(pre) {
				res = max(res, 1+dfs(i+1, cur))
			}
		}
		cache[pre] = res
		return res
	}
	return dfs(0, "")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestStrChainV2(words []string) int {
	f := make(map[string]int)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	ans := 0
	for _, word := range words {
		res := 0
		for i := range word {
			t := word[:i] + word[i+1:]
			res = max(res, 1+f[t])
		}
		f[word] = res
		ans = max(ans, res)
	}
	return ans
}
