package main

import "container/heap"

func main() {

}

// V2

type HP struct {
	data [][]int
}

func (h *HP) Len() int {
	return len(h.data)
}

func (h *HP) Less(i, j int) bool {
	return h.data[i][2] < h.data[j][2]
}

func (h *HP) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *HP) Push(v interface{}) {
	h.data = append(h.data, v.([]int))
}

func (h *HP) Pop() interface{} {
	v := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return v
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	hp := &HP{}
	/*
		关键解决思路：
			1、由于nums1和nums2都是升序排序，所以最小的一定是nums1[0] + nums2[0]
			2、第二小的一定是nums1[0] + nums2[1] 或者 nums1[1] + nums2[0]
			3、套到循环中：下一个最小值一定在 [i+1, j] 或者 [i, j+1] 中
			4、如果一开始堆中只放了[0, 0]，下一个入堆的是[0, 1]和[1, 0]，再下一个是[1,1],[0, 2],[1,1],[2,0]
				存在[1,1]重复的情况
			5、所以一开始把所有的 [i, 0] 都入堆，后续只增加 j 的值，也就是只把 [i, j+1] 入堆，保证不会重复
	*/
	for i, a := range nums1 {
		heap.Push(hp, []int{i, 0, a + nums2[0]})
		if hp.Len() >= k {
			break
		}
	}
	ans := make([][]int, 0, min(k, n*m))
	for hp.Len() > 0 && len(ans) < k {
		tp := heap.Pop(hp).([]int)
		i, j := tp[0], tp[1]
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < m {
			heap.Push(hp, []int{i, j + 1, nums1[i] + nums2[j+1]})
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
