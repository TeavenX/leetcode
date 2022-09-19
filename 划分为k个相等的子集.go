package main

import "sort"

func main() {

}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	bucket := make([]int, k)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var traceback func(idx int) bool
	traceback = func(idx int) bool {
		if idx == len(nums) {
			// 一定是满足的，不然回溯不到
			//for _, b := range bucket {
			//    if b != target {
			//        return false
			//    }
			//}
			return true
		}
		for i := range bucket {
			if i > 0 && bucket[i] == bucket[i-1] {
				continue
			}
			if bucket[i]+nums[idx] > target {
				continue
			}
			bucket[i] += nums[idx]
			if traceback(idx + 1) {
				return true
			}
			bucket[i] -= nums[idx]
		}
		return false
	}
	return traceback(0)
}
