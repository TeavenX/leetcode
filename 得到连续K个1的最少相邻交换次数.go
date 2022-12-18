package main

func main() {

}

func minMoves(nums []int, k int) int {
	genZeros := func() []int {
		result := make([]int, 0)
		left := -1
		for right, num := range nums {
			if num == 1 {
				if left != -1 {
					result = append(result, right-left-1)
				}
				left = right
			}
		}
		return result
	}
	zeros := genZeros()
	genPreSum := func() []int {
		result := []int{0}
		for i, num := range zeros {
			result = append(result, result[i]+num)
		}
		return result
	}
	preSum := genPreSum()
	genPreRange := func(i, j int) int {
		return preSum[j+1] - preSum[i]
	}
	cost := 0
	left, right := 0, k-2
	for i := left; i <= right; i++ {
		cost += zeros[i] * min(i+1, right-i+1)
	}
	minCost := cost
	for i, j := 1, 1+k-2; j < len(zeros); i++ {
		mid := (i + j) >> 1
		cost -= genPreRange(i-1, mid-1)
		cost += genPreRange(mid+k%2, j)
		minCost = min(cost, minCost)
		j++
	}
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
