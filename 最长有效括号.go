package main

func main() {

}

func longestValidParentheses(s string) int {
	n := len(s)
	f := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		if i > 0 && i-f[i-1] > 0 && s[i-f[i-1]-1] == '(' {
			f[i] = f[i-1] + 2
			if i-f[i-1]-2 > 0 {
				f[i] += f[i-f[i-1]-2]
			}
			ans = max(ans, f[i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
