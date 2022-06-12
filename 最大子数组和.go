package main

import "math"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArray(nums)
}

func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for _, num := range nums[1:] {
		newSum := sum + num
		if newSum > num {
			sum = newSum
		} else {
			sum = num
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArray20220506(nums []int) int {
	sum, max := nums[0], nums[0]
	for _, num := range nums[1:] {
		newSum := sum + num
		if newSum > num {
			// sum > 0
			sum = newSum
		} else {
			sum = num
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArrayDivide(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	crossMid := func(left, mid, right int) int {
		sumLeft := math.MinInt
		sum := 0
		for i := mid; i >= left; i-- {
			sum += nums[i]
			if sum > sumLeft {
				sumLeft = sum
			}
		}
		sumRight := math.MinInt
		sum = 0
		for i := mid + 1; i <= right; i++ {
			sum += nums[i]
			if sum > sumRight {
				sumRight = sum
			}
		}
		return sumLeft + sumRight
	}
	var subArraySum func(left, right int) int
	subArraySum = func(left, right int) int {
		if left == right {
			return nums[left]
		}
		mid := left + (right-left)>>1
		return max(max(subArraySum(left, mid), subArraySum(mid+1, right)), crossMid(left, mid, right))
	}
	return subArraySum(0, n-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
