package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if nums2[i] > nums1[j] {
			nums1[j+i+1] = nums2[i]
			i--
		} else {
			nums1[j+i+1] = nums1[j]
			j--
		}
	}
	if j < 0 {
		j = 0
	}
	for ; i >= 0; i-- {
		nums1[j+i] = nums2[i]
	}
}
