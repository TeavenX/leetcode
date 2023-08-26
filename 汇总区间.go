package main

import "fmt"

func main() {

}

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	n := len(nums)
	if n == 0 {
		return ans
	}
	start, end := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] != end+1 {
			ans = append(ans, embed(start, end))
			start = nums[i]
		}
		end = nums[i]
	}
	ans = append(ans, embed(start, end))
	return ans
}

func embed(start, end int) string {
	if start == end {
		return fmt.Sprintf("%d", start)
	}
	return fmt.Sprintf("%d->%d", start, end)
}
