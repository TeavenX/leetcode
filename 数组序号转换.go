package main

import "sort"

func main() {

}

func arrayRankTransform(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	sort.Ints(result)
	cache := make(map[int]int)
	i := 1
	for idx, num := range result {
		if idx > 0 && result[idx] == result[idx-1] {
			continue
		}
		cache[num] = i
		i++
	}
	for idx, num := range arr {
		result[idx] = cache[num]
	}
	return result
}
