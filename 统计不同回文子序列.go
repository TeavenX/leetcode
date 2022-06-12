package main

import "fmt"

func main() {
	//fmt.Println(countPalindromicSubsequences("abbad"))
	fmt.Println(countPalindromicSubsequences("bddaabdbbccdcdcbbdbddccbaaccabbcacbadbdadbccddccdbdbdbdabdbddcccadddaaddbcbcbabdcaccaacabdbdaccbaacc"))
}

func countPalindromicSubsequences(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for size := 2; size <= n; size++ {
		for i := 0; i <= n-size; i++ {
			j := size + i - 1
			if s[i] != s[j] {
				dp[i][j] = dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1]
			} else {
				dp[i][j] = 2 * dp[i+1][j-1]
				left, right := i+1, j-1
				for left <= right && s[left] != s[i] {
					left++
				}
				for left <= right && s[right] != s[j] {
					right--
				}
				if left > right {
					dp[i][j] += 2
				} else if left == right {
					dp[i][j] += 1
				} else {
					dp[i][j] -= dp[left+1][right-1]
				}
			}
			if dp[i][j] < 0 {
				dp[i][j] += 1e9 + 7
			}
			dp[i][j] %= 1e9 + 7
		}
	}
	return dp[0][n-1]
}
