package main

func main() {

}

func trap(height []int) int {
	n := len(height)
	premax, sufmax := make([]int, n), make([]int, n)
	premax[0] = height[0]
	sufmax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		premax[i] = max(premax[i-1], height[i])
		sufmax[n-1-i] = max(sufmax[n-i], height[n-1-i])
	}
	ans := 0
	for i, h := range height {
		ans += max(0, min(premax[i], sufmax[i])-h)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	n := len(height)
	left, right := 0, n-1
	lm := height[0]
	rm := height[n-1]
	ans := 0
	for left < right {
		if lm < rm {
			ans += lm - height[left]
			left++
			lm = max(lm, height[left])
		} else {
			ans += rm - height[right]
			right--
			rm = max(rm, height[right])
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
