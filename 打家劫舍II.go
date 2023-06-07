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

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(do(nums[1:]), do(nums[:n-1]))
}

func do(nums []int) int {
	res, pp, p := 0, 0, nums[0]
	for i := 1; i < len(nums); i++ {
		res = max(p, pp+nums[i])
		pp = p
		p = res
	}
	return p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
