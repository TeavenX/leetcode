package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	a := make([]int, 0)
	a = []int{-1, 0, 1, 2, -1, -4}
	a = []int{0, 0, 0}
	result := threeSumV3(a)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	cache := make(map[string]bool)
	result := make([][]int, 0)
	for i, num := range nums {
		for j := i + 1; j < len(nums)-1; j++ {
			num2 := nums[j]
			for k := j + 1; k < len(nums); k++ {
				num3 := nums[k]
				cacheKey := strconv.Itoa(num) + strconv.Itoa(num3) + strconv.Itoa(num3)
				if num+num2+num3 == 0 && !cache[cacheKey] {
					result = append(result, []int{num, num2, num3})
					cache[cacheKey] = true
				}
			}
		}
	}
	return result
}

func threeSumV2(nums []int) [][]int {
	sort.Ints(nums)
	cache := make(map[string]bool)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				cacheKey := strconv.Itoa(nums[i]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
				if !cache[cacheKey] {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					cache[cacheKey] = true
				}
				right--
				left++
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func threeSumV3(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
				for nums[right] == nums[right+1] && right > left {
					right--
				}
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
