package main

func main() {

}

func minOperations(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > 6*n2 || n2 > 6*n1 {
		return -1
	}
	s1, s2 := 0, 0
	for _, num := range nums1 {
		s1 += num
	}
	for _, num := range nums2 {
		s2 += num
	}
	if s1 < s2 {
		return core(nums2, nums1, s2, s1)
	}
	return core(nums1, nums2, s1, s2)
}

func core(nums1, nums2 []int, s1, s2 int) int {
	freq := [6]int{}
	for _, num := range nums1 {
		// nums1 可以减少 num-1
		freq[num-1]++
	}
	for _, num := range nums2 {
		// nums2 可以增加 6-num
		freq[6-num]++
	}
	diff := s1 - s2
	count := 0
	for i := 5; i >= 1 && diff > 0; i-- {
		// 两个相互靠近，先把大的数字处理了（贪心）
		for freq[i] > 0 && diff > 0 {
			count++
			freq[i]--
			diff -= i
		}
	}
	return count
}
