package main

func main() {

}

var steps = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func robotSim(commands []int, obstacles [][]int) int {
	x, y := 0, 0
	idx := 0
	stp := make(map[[2]int]bool)
	for _, ob := range obstacles {
		stp[[2]int{ob[0], ob[1]}] = true
	}
	valid := func(x, y int) bool {
		return !stp[[2]int{x, y}]
	}
	ans := 0
	for _, cmd := range commands {
		if cmd == -1 {
			idx = (idx + 4 + 1) % 4
		} else if cmd == -2 {
			idx = (idx + 4 - 1) % 4
		} else {
			for cmd > 0 {
				cmd--
				step := steps[idx]
				nx, ny := x+step[0], y+step[1]
				if !valid(nx, ny) {
					continue
				}
				x, y = nx, ny
			}
			ans = max(ans, x*x+y*y)
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
