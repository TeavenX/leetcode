package main

func main() {

}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robCore(nums[1:]), robCore(nums[:n-1]))
}

func robCore(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	pre, cur := 0, 0
	for _, num := range nums {
		temp := cur
		cur = max(cur, pre+num)
		pre = temp
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
