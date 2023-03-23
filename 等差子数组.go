package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 6, 5, 9, 3, 7}
	// l := []int{0, 0, 2}
	// r := []int{2, 3, 5}
	// fmt.Println(checkArithmeticSubarrays(nums, l, r))
	fmt.Println(nums[2:5])
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n := len(l)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		sub := append([]int{}, nums[l[i]:r[i]+1]...) // 这里会不用append会改变原数组的顺序
		sort.Ints(sub)
		ans[i] = true
		for j := 2; j < len(sub); j++ {
			if sub[j]-sub[j-1] != sub[1]-sub[0] {
				ans[i] = false
			}
		}
	}
	return ans
}
