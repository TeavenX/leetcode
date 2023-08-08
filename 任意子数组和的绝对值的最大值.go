package main

func main() {

}

func maxAbsoluteSum(nums []int) int {
	var sum, mx, mn int
	for _, num := range nums {
		sum += num
		if sum > mx {
			mx = sum
		} else if sum < mn {
			mn = sum
		}
	}
	return mx - mn
}

func maxAbsoluteSumV2(nums []int) int {
	var mx, mn, ans int
	for _, x := range nums {
		mx = max(mx, 0) + x
		mn = min(mn, 0) + x
		ans = max(ans, max(mx, -mn))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
