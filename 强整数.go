package main

func main() {

}

type Pair struct {
	x, y int
}

func powerfulIntegers(x int, y int, bound int) []int {
	if bound <= 1 {
		return []int{}
	}
	cache := make(map[Pair]struct{})
	var dfs func(x1, y1 int)
	dfs = func(x1, y1 int) {
		if x1+y1 > bound {
			return
		}
		key := Pair{x1, y1}
		if _, ok := cache[key]; ok {
			return
		}
		cache[key] = struct{}{}
		if v := x1*x + y1; v <= bound {
			dfs(x1*x, y1)
		}
		if v := x1 + y1*y; v <= bound {
			dfs(x1, y1*y)
		}
	}
	dfs(1, 1)
	ans := make([]int, 0, len(cache))
	m := make(map[int]bool)
	for k := range cache {
		if v := k.x + k.y; !m[v] {
			m[v] = true
			ans = append(ans, v)
		}
	}
	return ans
}
