package main

import "fmt"

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	s = "leetcode"
	wordDict = []string{"leet", "code"}
	s = "ab"
	wordDict = []string{"a", "b"}
	fmt.Println(wordBreak(s, wordDict))
}
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dict := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}
	if exist := dict[s]; exist {
		return true
	}
	dp := make([]bool, n)
	for i := 0; i < n; i++ {
		if exist := dict[s[:i+1]]; exist {
			dp[i] = true
		} else {
			for j := i; j > 0; j-- {
				if exist := dict[s[j:i+1]]; exist && dp[j-1] {
					dp[i] = true
				}
			}
		}
	}
	return dp[n-1]
}
