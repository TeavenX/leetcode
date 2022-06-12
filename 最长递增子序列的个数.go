package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findNumberOfLIS(nums))
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dpnum := make([]int, n)
	max := 1
	for i, num := range nums {
		dp[i] = 1
		dpnum[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < num {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					dpnum[i] = dpnum[j]
				} else if dp[j]+1 == dp[i] {
					dpnum[i] += dpnum[j]
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	result := 0
	for idx, num := range dp {
		if num == max {
			result += dpnum[idx]
		}
	}
	return result
}
