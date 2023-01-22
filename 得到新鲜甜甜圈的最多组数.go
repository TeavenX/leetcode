package main

func main() {

}

func maxHappyGroupsTLE(batchSize int, groups []int) int {
	cnt := make([]int, batchSize)
	for _, group := range groups {
		cnt[group%batchSize]++
	}
	var dfs func(left int, cnt []int) int
	dfs = func(left int, cnt []int) int {
		res := 0
		lc := 0
		if left == 0 {
			lc++
		}
		for i := range cnt {
			if cnt[i] > 0 {
				cnt[i]--
				res = max(res, lc+dfs((left+i+1)%batchSize, cnt))
				cnt[i]++
			}
		}
		return res
	}
	return cnt[0] + dfs(0, cnt[1:])
}

type pair struct {
	left, mask int
}

func maxHappyGroups(batchSize int, groups []int) int {
	ans := 0
	cnt := make([]int, batchSize)
	for _, g := range groups {
		g %= batchSize
		if g == 0 {
			ans++
		} else {
			c := batchSize - g
			if cnt[c] > 0 {
				cnt[c]--
				ans++
			} else {
				cnt[g]++
			}
		}
	}
	val := make([]int, 0)
	mask := 0

	for k, v := range cnt {
		if v > 0 {
			val = append(val, k)
			mask = mask<<5 | v
		}
	}

	for i, n := 0, len(val); i < n/2; i++ {
		val[i], val[n-i-1] = val[n-i-1], val[i]
	}

	cache := make(map[pair]int)
	var dfs func(left int, mask int) int
	dfs = func(left int, mask int) int {
		p := pair{left, mask}
		if c, exists := cache[p]; exists {
			return c
		}
		res := 0
		for i, c := range val {
			i *= 5
			if mask>>i&31 > 0 {
				r := dfs((left+c)%batchSize, mask-1<<i)
				if left == 0 {
					r++
				}
				res = max(res, r)
			}
		}
		cache[p] = res
		return res
	}
	return ans + dfs(0, mask)
}

func maxHappyGroupsV2(batchSize int, groups []int) int {
	ans := 0
	cnt := make([]int, batchSize)
	for _, g := range groups {
		g %= batchSize
		if g == 0 {
			ans++
		} else {
			c := batchSize - g
			if cnt[c] > 0 {
				cnt[c]--
				ans++
			} else {
				cnt[g]++
			}
		}
	}

	val := make([]int, 0)
	mask := 0
	for i, c := range cnt {
		if c > 0 {
			val = append(val, i)
		}
		mask = mask<<5 | c
	}

	n := len(val)

	for i := 0; i < n/2; i++ {
		val[i], val[n-1-i] = val[n-1-i], val[i]
	}

	cache := make(map[int]int)
	var dfs func(mask int) int
	dfs = func(mask int) int {
		if v, exists := cache[mask]; exists {
			return v
		}
		res := 0
		left := mask >> 20
		for i, v := range val {
			i *= 5
			if mask>>i&31 > 0 {
				// todo
				r := dfs((((left + v) % batchSize) << 20) | mask - (1 << i))
				if left == 0 {
					r++
				}
				res = max(res, r)
			}
		}
		cache[mask] = res
		return res
	}
	return ans + dfs(mask)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
