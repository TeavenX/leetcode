package main

func main() {

}

func splitArraySameAverage(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	n := len(nums)
	m := n >> 1

	possible := false
	for i := 1; i <= m; i++ {
		if sum*i%n == 0 {
			possible = true
			break
		}
	}
	if !possible {
		return false
	}

	dp := make([]map[int]bool, m+1)
	for i := range dp {
		dp[i] = map[int]bool{}
	}
	dp[0][0] = true
	for _, num := range nums {
		for i := m; i >= 1; i-- {
			for x := range dp[i-1] {
				curr := x + num
				if curr*n == sum*i {
					return true
				}
				dp[i][curr] = true
			}
		}
	}
	return false
}
