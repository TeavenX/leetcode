package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums = []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	max := 1
	for i, num := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < num && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
