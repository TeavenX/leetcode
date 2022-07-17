package main

func main() {

}

func arrayNesting(nums []int) int {
	n := len(nums)
	maxCount := 0
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			continue
		}
		t, count := i, 0
		for t < n && nums[t] >= 0 {
			count++
			temp := nums[t]
			nums[t] = -1
			t = temp
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func arrayNestingV2(nums []int) int {
	n := len(nums)
	maxCount := 0
	for i := 0; i < n; i++ {
		count := 0
		for nums[i] < n {
			i, nums[i] = nums[i], n
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
