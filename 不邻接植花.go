package main

func main() {

}

func gardenNoAdj(n int, paths [][]int) []int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0, 3)
	}
	for _, path := range paths {
		a, b := path[0]-1, path[1]-1
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		use := 0
		for _, node := range g[i] {
			use |= 1 << ans[node]
		}
		use ^= 0b11110
		for j := 1; j <= 4; j++ {
			if use>>j&1 == 0 {
				continue
			}
			ans[i] = j
			break
		}
	}
	return ans
}
