package main

import "sort"

func main() {

}

func canBeEqual(target []int, arr []int) bool {
	cache := make(map[int]int)
	for _, num := range target {
		cache[num]++
	}
	for _, num := range arr {
		if cache[num] == 0 {
			return false
		}
		cache[num]--
	}
	return true
}

func canBeEqualV2(target []int, arr []int) bool {
	cache := [1001]int{}
	for _, num := range target {
		cache[num]++
	}
	for _, num := range arr {
		if cache[num] == 0 {
			return false
		}
		cache[num]--
	}
	return true
}

func canBeEqualV3(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i, num := range target {
		if num != arr[i] {
			return false
		}
	}
	return true
}
