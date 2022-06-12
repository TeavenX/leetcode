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

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
	}
	p1, p2 := m-1, n-1
	for p2 >= 0 && p1 >= 0 {
		if nums2[p2] >= nums1[p1] {
			nums1[p2+p1+1] = nums2[p2]
			p2--
		} else {
			nums1[p2+p1+1] = nums1[p1]
			p1--
		}
	}
	for p2 >= 0 {
		nums1[p2] = nums2[p2]
		p2--
	}
}
