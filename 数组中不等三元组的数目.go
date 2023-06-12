package main

import "sort"

func main() {

}

func unequalTriplets(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	start := 0
	for i := 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			ans += start * (i - start + 1) * (n - i - 1)
			start = i + 1
		}
	}
	return ans
}
