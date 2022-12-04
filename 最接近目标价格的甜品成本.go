package main

import (
	"fmt"
	"math"
)

func main() {
	baseCosts := []int{3, 10}
	toppingCosts := []int{2, 5}
	fmt.Println(closestCost(baseCosts, toppingCosts, 9))
}

func closestCostV2(baseCosts []int, toppingCosts []int, target int) int {
	base := baseCosts[0]
	for _, c := range baseCosts {
		base = min(base, c)
	}
	if base > target {
		return base
	}
	can := make([]bool, target+1)    // 没有理解
	result := target + target - base // 没有理解
	for _, c := range baseCosts {
		if c <= target {
			can[c] = true
		} else {
			result = min(result, c) // 没有理解
		}
	}
	for _, c := range toppingCosts {
		for count := 0; count < 2; count++ { // 每种配料能选两次, 为什么放在第二层而不是第三层
			for i := target; i > 0; i-- {
				if can[i] && i+c > target {
					result = min(result, i+c)
				}
				if i-c > 0 {
					can[i] = can[i] || can[i-c]
				}
			}
		}
	}
	for i := 0; i <= result-target; i++ {
		if can[target-i] {
			return target - i
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	result := math.MaxInt
	n := len(toppingCosts)
	var dfs func(base int, topidx int)
	dfs = func(base int, topidx int) {
		a := abs(target - base)
		b := abs(target - result)
		if a < b || (a == b && base < result) {
			result = base
		}
		if base > target {
			return
		}
		for i := topidx; i < n; i++ {
			dfs(base+toppingCosts[i], i+1)
			dfs(base+2*toppingCosts[i], i+1)
		}
	}
	for _, c := range baseCosts {
		dfs(c, 0)
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
