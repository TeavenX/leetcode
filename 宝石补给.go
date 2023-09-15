package main

func main() {

}

func giveGem(gem []int, operations [][]int) int {
	for _, op := range operations {
		t := gem[op[0]] >> 1
		gem[op[0]] -= t
		gem[op[1]] += t
	}
	mx, mn := gem[0], gem[0]
	for _, g := range gem {
		mx = max(mx, g)
		mn = min(mn, g)
	}
	return mx - mn
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
