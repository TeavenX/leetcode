package main

import (
	"fmt"
	"math"
)

func main() {
	stones := []int{3, 2, 4, 1}
	fmt.Println(mergeStones(stones, 2))
}

func mergeStonesErr(stones []int, k int) int {
	n := len(stones)
	if n == 0 {
		return 0
	} else if n == 1 {
		return stones[0]
	} else if n < k {
		return -1
	}
	ans := math.MaxInt32
	for i := 0; i <= n-k; i++ {
		pre := stones[:i]
		suf := stones[i+k:]
		mid := 0
		for j := i; j < i+k; j++ {
			mid += stones[j]
		}
		arr := append(append(append([]int{}, pre...), mid), suf...)
		res := mergeStones(arr, k)
		if res < 0 {
			continue
		}
		ans = min(ans, mid+res)
	}
	return ans
}

func mergeStones(stones []int, k int) int {
	n := len(stones)

	if (n-1)%(k-1) != 0 {
		// n堆合并到1堆，需要减少n-1堆
		// k堆合并到1堆减少了k-1堆
		// 也就是需要合并 (n-1)/(k-1)次
		return -1
	}

	s := make([]int, n+1)
	for i, num := range stones {
		s[i+1] = s[i] + num
	}

	cache := make([][][]int, n)
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, k+1)
			for idx := range cache[i][j] {
				cache[i][j][idx] = -1
			}
		}
	}
	// 将左边界 left 到右边界 right 合并为 target 堆的成本
	var dfs func(left, right, target int) int
	dfs = func(left, right, target int) (ans int) {
		if v := cache[left][right][target]; v > 0 {
			return v
		}
		defer func() {
			cache[left][right][target] = ans
		}()
		if target == 1 {
			if left == right {
				return 0
			}
			// target=1，也就是需要将 K 堆合并成一堆
			return dfs(left, right, k) + s[right+1] - s[left]
		}
		ans = math.MaxInt
		for i := left; i < right; i += k - 1 {
			// 这里只需要处理 1，k-1 的原因是 dfs 到后面会包含 x，k-x的场景
			ans = min(ans, dfs(left, i, 1)+dfs(i+1, right, target-1))
		}
		return ans
	}
	return dfs(0, n-1, 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
