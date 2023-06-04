package main

func main() {

}

func applyOperations(nums []int) []int {
	n := len(nums)
	ans := nums[:0] // 这里会复用nums的内存空间
	for i := 0; i < n-1; i++ {
		if nums[i] > 0 {
			if nums[i] == nums[i+1] {
				ans = append(ans, 2*nums[i])
				nums[i+1] = 0
			} else {
				ans = append(ans, nums[i])
			}
		}
	}
	ans = append(ans, nums[n-1])
	for len(ans) < n {
		ans = append(ans, 0)
	}
	return ans
}
