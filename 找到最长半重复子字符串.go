package main

import "fmt"

func main() {
	// fmt.Println(longestSemiRepetitiveSubstring("0111"))
	// fmt.Println(longestSemiRepetitiveSubstring("52233"))
	fmt.Println(longestSemiRepetitiveSubstring("1000"))
	// fmt.Println(longestSemiRepetitiveSubstring("1111111"))
}

func longestSemiRepetitiveSubstring(s string) int {
	ans := 1
	n := len(s)
	same := 0
	left := 0
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			same++
			if same > 1 {
				left++
				// 因为假设字符串是 123445678890
				// 遍历到第二个8的时候same大于1了
				// 此时left = 0，中间包含了44，需要把44去掉
				// 所以这里要一直增长left，直到遇到第一个重复的，此时left = index(第二个4)
				// 如果是 111111 的情况
				// 遍历到第三个1的时候，same>1
				// s[left] == s[left-1]
				for s[left] != s[left-1] {
					left++
				}
				same = 1
			}
		}
		ans = max(ans, i-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
