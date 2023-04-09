package main

import "math/bits"

func main() {

}

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	skillMap := make(map[string]int)
	want := 0
	for i, skill := range req_skills {
		skillMap[skill] = i
		want |= 1 << i
	}
	n := len(people)
	peopleCache := make([]int, n)
	var all int64 = 0
	for i, p := range people {
		all |= 1 << i
		for _, skill := range p {
			peopleCache[i] |= 1 << skillMap[skill]
		}
	}
	cache := make([][]int64, n)
	for i := range cache {
		cache[i] = make([]int64, want+1)
	}
	var dfs func(want int, idx int) int64
	dfs = func(want int, idx int) (ans int64) {
		if want == 0 {
			return 0
		}
		if idx == n {
			return all
		}
		if v := cache[idx][want]; v > 0 {
			return v
		}
		defer func() {
			cache[idx][want] = ans
		}()
		ans = all
		for i := idx; i < n; i++ {
			if peopleCache[i] == 0 {
				continue
			}
			t := want ^ peopleCache[i]
			res := dfs(want&t, i+1)
			res |= 1 << i
			if bits.OnesCount64(uint64(ans)) > bits.OnesCount64(uint64(res)) {
				ans = res
			}
		}
		return
	}
	res := dfs(want, 0)
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if res>>i&1 == 1 {
			ans = append(ans, i)
		}
	}
	return ans
}
