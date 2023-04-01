package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(arithmeticTriplets(nums, 2))
}

func arithmeticTripletsErr(nums []int, diff int) int {
	ans := 0
	n := len(nums)
	for k := n - 1; k >= 0; k-- {
		for i := 0; i < k-1; i++ {
			left, right := i+1, k-1
			for left <= right {
				j := (left + right) >> 1
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					ans++
					left++
					right--
					fmt.Println(nums[i], nums[j], nums[k])
				} else if nums[j]-nums[i] > diff || nums[k]-nums[j] < diff {
					right = j - 1
				} else {
					left = j + 1
				}
			}
		}
	}
	return ans
}

func arithmeticTriplets(nums []int, diff int) int {
	ans := 0
	n := len(nums)

	search := func(left, right, target int) (idx int, found bool) {
		for left <= right {
			mid := (left + right) >> 1
			if nums[mid] == target {
				return mid, true
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1, false
	}

	for i, num := range nums {
		target := num + diff
		j, found := search(i+1, n-2, target)
		if !found {
			continue
		}
		target += diff
		_, found = search(j+1, n-1, target)
		if found {
			ans++
		}
	}
	return ans
}
