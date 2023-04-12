package main

import "sort"

func main() {

}

func mostFrequentEven(nums []int) int {
	sort.Ints(nums)
	count := 0
	max := -1
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 0 {
			c := 1
			for i < len(nums)-1 && nums[i+1] == nums[i] {
				c++
				i++
			}
			if c > count {
				max = nums[i]
				count = c
			}
		}
	}
	return max
}
