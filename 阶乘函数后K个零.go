package main

import "math"

func main() {

}

func preimageSizeFZF(k int) int {
	left, right := 0, k
	for left <= right {
		mid := (right-left)>>1 + left
		count := mid
		// 上界是14的原因是题目给予k的范围是[0,10^9], 5^13 > 10^9 而 5^12 < 10^9
		for i := 1; i < 14; i++ {
			count += mid / int(math.Pow(float64(5), float64(i)))
		}
		if count == k {
			return 5
		}
		if count < k {
			left = mid + 1
		} else if count > k {
			right = mid - 1
		}
	}
	return 0
}
