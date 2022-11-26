package main

func main() {

}

func check(nums []int) bool {
	diffIdx := 0
	min := 100
	for i, num := range nums {
		if i > 0 && num < nums[i-1] {
			if diffIdx != 0 {
				return false
			}
			diffIdx = i
			if min < nums[len(nums)-1] {
				return false
			}
		}
		if num < min {
			min = num
		}
	}
	return true
}
