package main

func main() {

}

func partitionDisjoint(nums []int) int {
	max := nums[0]
	leftMax := nums[0]
	idx := 0
	for i, num := range nums {
		if num > max {
			max = num
		}
		if num < leftMax {
			leftMax = max
			idx = i
		}
	}
	return idx + 1
}
