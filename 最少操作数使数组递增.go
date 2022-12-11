package main

func main() {

}

func minOperations(nums []int) int {
	count := 0
	pre := nums[0]
	for i, num := range nums {
		if i > 0 {
			if num < pre+1 {
				count += pre + 1 - num
				num = pre + 1
			}
			pre = num
		}
	}
	return count
}
