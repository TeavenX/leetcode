package main

func main() {

}

func countServers(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	row := make([]int, m)
	col := make([]int, n)
	ans := 0
	for j, r := range grid {
		for i, v := range r {
			if v > 0 {
				row[i]++
				col[j]++
			}
		}
	}
	for j, r := range grid {
		for i, v := range r {
			if v == 1 && (row[i] > 1 || col[j] > 1) {
				ans++
			}
		}
	}
	return ans
}
