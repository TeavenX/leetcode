package main

import "math/rand"

func main() {

}

func findKthLargest(nums []int, k int) int {
	partition := func(left, right int) int {
		pos := left + rand.Intn(right-left+1)
		nums[left], nums[pos] = nums[pos], nums[left]

		pivot := nums[left]

		le, ge := left+1, right
		for {
			for le <= ge && nums[le] < pivot {
				le++
			}
			for le <= ge && nums[ge] > pivot {
				ge--
			}
			if le >= ge {
				break
			}
			nums[le], nums[ge] = nums[ge], nums[le]
			le++
			ge--
		}
		nums[left], nums[ge] = nums[ge], nums[left]
		return ge
	}
	n := len(nums)
	target := n - k
	l, r := 0, n-1
	for {
		idx := partition(l, r)
		if idx == target {
			return nums[idx]
		}
		if idx < target {
			l = idx + 1
		} else {
			r = idx - 1
		}
	}
}
