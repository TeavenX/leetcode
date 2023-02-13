package main

func main() {

}

func balancedString(s string) int {
	cache := ['X']int{}
	avg := len(s) >> 2
	for _, b := range s {
		cache[b]++
	}
	if cache['Q'] == avg && cache['W'] == avg && cache['E'] == avg && cache['R'] == avg {
		return 0
	}
	ans, left := len(s), 0
	for right, b := range s {
		cache[b]--
		for cache['Q'] <= avg && cache['W'] <= avg && cache['E'] <= avg && cache['R'] <= avg {
			ans = min(ans, right-left+1)
			cache[s[left]]++
			left++
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
