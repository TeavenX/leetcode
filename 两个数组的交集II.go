package main

import "sort"

func main() {

}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	p1, p2 := 0, 0
	result := make([]int, 0)
	findEqual := func() {
		if nums1[p1] > nums2[p2] {
			nums1, nums2 = nums2, nums1
			p1, p2 = p2, p1
		}
		for p1 < len(nums1) && p2 < len(nums2) && nums1[p1] < nums2[p2] {
			p1++
		}
		if p1 >= len(nums1) || p2 >= len(nums2) {
			return
		}
		for p1 < len(nums1) && p2 < len(nums2) && nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++
		}
	}
	for p1 < len(nums1) && p2 < len(nums2) {
		findEqual()
	}
	return result
}
