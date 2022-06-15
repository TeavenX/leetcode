package main

import "sort"

func main() {

}

/*
最优解一定是相邻的三个元素
有没有可能，最优解是不相邻的三个元素，「没有」
假设最优解是a,b,c，a < b < c，此时令ab之间存在a', a < a' < b，bc之间存在b'，b < b' < c
根据三角形三边原则，由于 a + b > c，而 a < a'，那么a' + b > c; b + c > a，因为b > a'，所以 b + c > a'恒成立
因此a'不存在，同理b'也不存在，所以最优解一定是相邻的三个元素
*/
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}
