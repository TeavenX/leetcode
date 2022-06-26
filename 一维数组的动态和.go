package main

func main() {

}

func runningSum(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return []int{nums[0]}
	}
	result := make([]int, n)
	result[0] = nums[0]
	for i, num := range nums[1:] {
		result[i+1] = result[i] + num
	}
	return result
}

func runningSumV2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
