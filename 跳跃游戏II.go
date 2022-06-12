package main

func main() {

}

func jump(nums []int) int {
	end, count, maxLen := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxLen = max(nums[i]+i, maxLen)
		if i == end {
			count++
			end = maxLen
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
