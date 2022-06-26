package main

func main() {

}

func pivotIndex(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	if nums[n-1]-nums[0] == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		if nums[i-1] == nums[n-1]-nums[i] {
			return i
		}
	}
	return -1
}

func pivotIndexV2(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, num := range nums {
		if sum*2+num == total {
			return i
		}
		sum += num
	}
	return -1
}
