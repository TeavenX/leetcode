package main

func main() {

}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, num := range nums {
		s[i+1] = s[i] + num
	}
	search := func(l, r int) int {
		sum := 0
		for i := 0; i <= n-l-r; i++ {
			sr := 0
			for j := i + l; j <= n-r; j++ {
				sr = max(sr, s[j+r]-s[j])
			}
			sum = max(sum, s[i+l]-s[i]+sr)
		}
		return sum
	}
	return max(search(firstLen, secondLen), search(secondLen, firstLen))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
