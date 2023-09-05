package main

func main() {

}

func minNumber(nums1 []int, nums2 []int) int {
	mn1, mn2 := nums1[0], nums2[0]
	cache := [10]bool{}
	for _, x := range nums1 {
		mn1 = min(mn1, x)
		cache[x] = true
	}
	ans := 100
	for _, x := range nums2 {
		mn2 = min(mn2, x)
		if cache[x] {
			ans = min(ans, x)
		}
	}
	return min(ans, min(mn1*10+mn2, mn2*10+mn1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
