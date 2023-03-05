package main

import "strings"

func main() {

}

func minimumDeletions(s string) int {
	del := strings.Count(s, "a")
	ans := del
	for _, b := range s {
		if b == 'a' {
			// 前面的a不用删，所以这里只需要记录后面有多少a就行
			del--
		} else {
			// 因为前面不能有b，所以碰到b都要记录删除，直到把所有b都删掉
			del++
		}
		if del < ans {
			ans = del
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
