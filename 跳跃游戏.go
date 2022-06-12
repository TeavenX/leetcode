package main

func main() {

}

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	for i := n - 1; i >= 0; i-- {
		if nums[i] > 0 || (dp[i] == false && nums[i] > 1) {
			dp[i] = true
		} else {
			dp[i] = false
		}
	}

}

func canJump(nums []int) bool {
	n := len(nums)
	zeroCount := 0
	for i := n - 1; i >= 0; i-- {
		if nums[i] > 0 || nums[i] > zeroCount {
			zeroCount = 0
		} else {
			zeroCount++
		}
	}
	return zeroCount == 0
}
