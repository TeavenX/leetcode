package main

import "fmt"

func main() {
	//fmt.Println(longestPalindromeV2("babad"))
	//fmt.Println(longestPalindromeV2("ccc"))
	//fmt.Println(longestPalindromeV2("aacabdkacaa"))
	//fmt.Println(longestPalindromeV2("cbbd"))
	fmt.Println(longestPalindromeV2("ac"))
}

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	dist := 1
	ri, rj := 0, 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if j-i < 3 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] {
				if j-i >= dist {
					dist = j - i
					ri, rj = i, j
				}
			}
		}
	}
	return s[ri : rj+1]
}

func longestPalindromeV2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	maxLen, begin := 0, 0
	for i := 0; i < n; i++ {
		left1, right1 := extend(s, n, i, i)
		left2, right2 := extend(s, n, i, i+1)
		if right1-left1 > maxLen {
			maxLen = right1 - left1
			begin = left1
		}
		if right2-left2 > maxLen {
			maxLen = right2 - left2
			begin = left2
		}
	}
	return s[begin : begin+maxLen+1]
}

func extend(s string, n, left, right int) (int, int) {
	for left >= 0 && right < n && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
