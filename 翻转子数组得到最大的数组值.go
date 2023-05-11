package main

func main() {

}

func maxValueAfterReverse(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	cur := 0
	for i := 0; i < n-1; i++ {
		cur += abs(nums[i] - nums[i+1])
	}
	res := cur
	// left side = 0
	for i := 1; i < n-1; i++ {
		res = max(res, cur+abs(nums[i+1]-nums[0])-abs(nums[i+1]-nums[i]))
	}
	// right side = n-1
	for i := 1; i < n-1; i++ {
		res = max(res, cur+abs(nums[i-1]-nums[n-1])-abs(nums[i]-nums[i-1]))
	}
	a, b := min(nums[0], nums[1]), max(nums[0], nums[1])
	for i := 0; i < n-1; i++ {
		a = max(a, min(nums[i], nums[i+1]))
		b = min(b, max(nums[i], nums[i+1]))
	}
	return max(res, cur+2*(a-b))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
