package main

func main() {

}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	cache := make(map[int]int)
	result := make([]int, 0)
	for i, nums := range [][]int{nums1, nums2, nums3} {
		for _, num := range nums {
			cache[num] |= 1 << i
		}
	}
	for k, v := range cache {
		if v&(v-1) > 0 {
			result = append(result, k)
		}
	}
	return result
}
