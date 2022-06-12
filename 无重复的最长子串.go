package main

import "fmt"

func main() {

	s := "abcabcaaa"
	s = "aaaaaaaaaa"
	s = "a"
	s = "bbcb"
	s = "bbcbab"
	s = "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	strMap := map[byte]int{}
	left, right, n := 0, 0, len(s)
	maxLen := 0
	for right < n {
		str := s[right]
		if newLeft, contain := strMap[str]; contain {
			if newLeft > left {
				left = newLeft
			}
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		strMap[str] = right + 1
		right++
	}
	return maxLen
}

func lengthOfLongestSubstring3(s string) int {
	var n = len(s)
	if n <= 1 {
		return n
	}
	var maxLen = 1
	var left, right, window = 0, 0, make(map[byte]int)
	for right < n {
		var rightChar = s[right]
		var rightCharIndex = 0
		if _, ok := window[rightChar]; ok {
			rightCharIndex = window[rightChar]
		}
		if rightCharIndex > left {
			left = rightCharIndex
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		window[rightChar] = right + 1
		right++
	}
	return maxLen
}
