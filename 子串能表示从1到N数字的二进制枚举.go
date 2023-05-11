package main

func main() {

}

func queryStringOOM(s string, n int) bool {
	st := 0
	for st < len(s) && s[st] == '0' {
		st++
	}
	s = s[st:]
	m := len(s)
	cache := make([]bool, n+1)
	var dfs func(i int, pre int)
	dfs = func(i int, pre int) {
		if i == m {
			return
		}
		pre = pre<<1 + int(s[i]-'0')
		if pre > n {
			return
		}
		cache[pre] = true
		dfs(i+1, pre)
	}
	for i := range s {
		dfs(i, 0)
	}
	for _, v := range cache[1:] {
		if !v {
			return false
		}
	}
	return true
}

func queryString(s string, n int) bool {
	m := len(s)
	cache := make(map[int]struct{})
	for i := 0; i < m; i++ {
		v := int(s[i] - '0')
		if v == 0 {
			continue
		}
		for j := i + 1; v <= n; j++ {
			cache[v] = struct{}{}
			if j == len(s) {
				break
			}
			v = v<<1 + int(s[j]-'0')
		}
	}
	return len(cache) == n
}
