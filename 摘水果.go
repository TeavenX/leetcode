package main

func main() {

}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	s := make([]int, n+1)
	for i, c := range fruits {
		s[i+1] = s[i] + c[1]
	}
	ans := 0
	left := 0
	for ; left < n && fruits[left][0] < startPos-k; left++ {
	}
	for right := left; right < n && fruits[right][0] <= startPos+k; right++ {
		for left <= right && right < n &&
			startPos-fruits[left][0]+fruits[right][0]-fruits[left][0] > k &&
			fruits[right][0]-startPos+fruits[right][0]-fruits[left][0] > k {
			left++
		}
		ans = max(ans, s[right+1]-s[left])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
