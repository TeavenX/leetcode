package main

func main() {

}

func maxSubarraySumCircular(nums []int) int {
	total := nums[0]
	curMax := nums[0]
	curMin := nums[0]
	sumMax := nums[0]
	sumMin := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		total += num
		curMax = max(curMax+num, num)
		curMin = min(curMin+num, num)
		sumMax = max(sumMax, curMax)
		sumMin = min(sumMin, curMin)
	}
	if total == sumMin {
		return sumMax
	}
	return max(sumMax, total-sumMin)
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
