package main

func main() {

}

func maxScore(nums []int) int {
	gcds := [15][15]int{}
	n := len(nums)
	for i, x := range nums {
		for j := i + 1; j < n; j++ {
			gcds[i][j] = gcd(x, nums[j])
		}
	}
	cache := make(map[int]int)
	var helper func(num, idx int) int
	helper = func(num, idx int) int {
		if num == 0 {
			return 0
		}
		// 最多操作7次，因此将后三位记录操作次数
		key := num<<3 + idx
		if r, ok := cache[key]; ok {
			return r
		}

		maxScore := 0
		for i := 0; i < n; i++ {
			if (num>>i)&1 == 0 {
				// 已经被选过了
				continue
			}
			for j := i + 1; j < n; j++ {
				if num>>j&1 == 0 {
					continue
				}
				maxScore = max(maxScore, idx*gcds[i][j]+helper(num^(1<<i)^(1<<j), idx+1))
			}
		}
		cache[key] = maxScore
		return maxScore
	}
	// 用n位代表哪些数字可以选，1可以选，0已经被选了
	return helper(1<<n-1, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
