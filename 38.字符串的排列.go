package main

import "strconv"

func main() {

}

func permutation(s string) []string {
	ans := make([]string, 0)
	n := len(s)
	cache := make(map[string]bool)
	var dfs func(pre string, idx int)
	dfs = func(pre string, idx int) {
		key := pre + ";" + strconv.Itoa(idx)
		if cache[key] {
			return
		}
		cache[key] = true
		if idx == 1<<n-1 {
			ans = append(ans, pre)
			return
		}
		for i := 0; i < n; i++ {
			if idx>>i&1 == 0 {
				dfs(pre+string(s[i]), idx|1<<i)
			}
		}
	}
	dfs("", 0)
	return ans
}

func permutation(s string) []string {
	ans := make([]string, 0)
	n := len(s)
	bt := make([]byte, 0, n)
	cache := make(map[string]struct{})
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == 1<<n-1 {
			cache[string(bt)] = struct{}{}
			return
		}
		for i := 0; i < n; i++ {
			if idx>>i&1 == 0 {
				bt = append(bt, s[i])
				dfs(idx | 1<<i)
				bt = bt[:len(bt)-1]
			}
		}
	}
	dfs(0)
	for k, _ := range cache {
		ans = append(ans, k)
	}
	return ans
}
