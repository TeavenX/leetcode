package main

func main() {

}

func minMaxGame(nums []int) int {
	n := len(nums)
	for n > 1 {
		for i := 0; i < n>>1; i++ {
			if i&1 == 0 {
				nums[i] = min(nums[i<<1], nums[i<<1+1])
			} else {
				nums[i] = max(nums[i<<1], nums[i<<1+1])
			}
		}
		n >>= 1
	}
	return nums[0]
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
