package main

func main() {

}

func minTaps(n int, ranges []int) int {
	cache := make([]int, n+1)
	for i, d := range ranges {
		l := max(i-d, 0)
		cache[l] = max(cache[l], i+d)
	}
	ans := 0
	curMax := 0
	nextMax := 0
	for i, d := range cache {
		nextMax = max(nextMax, d)
		if i == curMax {
			if i == nextMax && i < n {
				return -1
			} else if i >= n {
				return ans
			}
			curMax = max(curMax, nextMax)
			ans++
		}
	}
	return ans
}

func minTaps(n int, ranges []int) int {
	cache := make([]int, n+1)
	for i, d := range ranges {
		l := max(i-d, 0)
		cache[l] = max(cache[l], i+d)
	}
	ans := 0
	curMax := 0
	nextMax := 0
	for i, d := range cache[:n] {
		// 这里只遍历到n的原因是，如果i > nextMax了，那一定能走到n，就足够了（只需要覆盖n的左边）
		// 等同于上面的写法
		nextMax = max(nextMax, d)
		if i == curMax {
			if i == nextMax {
				return -1
			}
			curMax = max(curMax, nextMax)
			ans++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
