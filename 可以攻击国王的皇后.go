package main

func main() {

}

var forward = [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	cache := make(map[int]bool)
	for _, q := range queens {
		cache[q[0]*10+q[1]] = true
	}
	ans := make([][]int, 0)
	step := func(x, y int, i, j int) {
		for x >= 0 && y >= 0 && x < 8 && y < 8 {
			if cache[x*10+y] {
				ans = append(ans, []int{x, y})
				return
			}
			x += i
			y += j
		}
	}
	for _, fwd := range forward {
		step(king[0], king[1], fwd[0], fwd[1])
	}
	return ans
}
