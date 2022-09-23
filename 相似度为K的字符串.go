package main

import "fmt"

func main() {
	//fmt.Println(kSimilarity("abac", "baca"))
	fmt.Println(kSimilarity("abcdeabcdeabcdeabcde", "aaaabbbbccccddddeeee"))
}

func kSimilarityError(s1 string, s2 string) int {
	n := len(s1)
	cache := make(map[string]bool)
	result := n - 1
	var dfs func(s1, s2 string, idx, cost int)
	dfs = func(s1, s2 string, idx, cost int) {
		if cache[s2] {
			return
		}
		cache[s2] = true
		if cost > result {
			return
		}
		if idx == n || s1 == s2 {
			result = min(result, cost)
			return
		}
		for i := idx; i < n; i++ {
			if s1[i] == s2[i] {
				continue
			}
			b2 := []byte(s2)
			for j := i + 1; j < n; j++ {
				if s1[i] == s2[j] && s1[j] != s2[j] {
					b2[i], b2[j] = b2[j], b2[i]
					dfs(s1, string(b2), i+1, cost+1)
					b2[i], b2[j] = b2[j], b2[i]
				}
			}
		}
	}
	dfs(s1, s2, 0, 0)
	return result
}

func kSimilarity(s1 string, s2 string) int {
	s, t := make([]byte, 0), make([]byte, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			s = append(s, s1[i])
			t = append(t, s2[i])
		}
	}
	n := len(s)
	if n == 0 {
		return 0
	}
	result := n - 1
	var dfs func(idx, cost int)
	dfs = func(idx, cost int) {
		if cost > result {
			return
		}
		for idx < n && s[idx] == t[idx] {
			idx++
		}
		if idx == n {
			result = min(result, cost)
			return
		}
		for j := idx + 1; j < n; j++ {
			if s[j] == t[idx] {
				s[idx], s[j] = s[j], s[idx]
				dfs(idx+1, cost+1)
				s[idx], s[j] = s[j], s[idx]
			}
		}
	}
	dfs(0, 0)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
