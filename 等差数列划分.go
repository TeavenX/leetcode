package main

func main() {

}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	result := 0
	length := 0
	for i := 2; i < n; i++ {
		if nums[i] == nums[i-1]+nums[i-1]-nums[i-2] {
			length += 1
			result += length
		} else {
			length = 0
		}
	}
	return result
}
