package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countVowelStringsV3(2))
	fmt.Println(countVowelStringsV3(1))
	fmt.Println(countVowelStringsV3(33))
	fmt.Println(countVowelStringsV3(37))
}

func countVowelStringsTLE(n int) int {
	cache := make(map[string]int)
	var dfs func(num string, i int) int
	dfs = func(num string, i int) int {
		if i == n {
			return 1
		}
		key := strconv.Itoa(i) + num
		if v, ok := cache[key]; ok {
			return v
		}
		res := 0
		var start byte
		if i == 0 {
			start = 1
		} else {
			start = num[i-1] - '0'
		}
		for j := start; j <= 5; j++ {
			res += dfs(num[:i]+string(j+'0')+num[i+1:], i+1)
		}
		cache[key] = res
		return res
	}
	num := "1"
	for i := 1; i < n; i++ {
		num += "1"
	}
	return dfs(num, 0)
}

func countVowelStringsV2(n int) int {
	cache := make([][6]int, n)
	var dfs func(pre int, i int) int
	dfs = func(pre int, i int) int {
		if i == n {
			return 1
		}
		if v := cache[i][pre]; v > 0 {
			return v
		}
		res := 0
		for j := pre; j <= 5; j++ {
			res += dfs(j, i+1)
		}
		cache[i][pre] = res
		return res
	}
	return dfs(1, 0)
}

func countVowelStringsV3(n int) int {
	dp := [...]int{1, 1, 1, 1, 1}
	// for i := 0; i < n; i++ {
	// 	dp[0] = dp[0] + dp[1] + dp[2] + dp[3] + dp[4]
	// 	dp[1] = dp[1] + dp[2] + dp[3] + dp[4]
	// 	dp[2] = dp[2] + dp[3] + dp[4]
	// 	dp[3] = dp[3] + dp[4]
	// }
	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				dp[j] += dp[k]
			}
		}
	}
	return dp[0]
}
