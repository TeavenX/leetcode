package main

func main() {

}

func oddCells(m int, n int, indices [][]int) int {
	col, row := make([]int, n), make([]int, m)
	for _, ind := range indices {
		row[ind[0]]++
		col[ind[1]]++
	}
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (row[i]+col[j])&1 == 1 {
				result++
			}
		}
	}
	return result
}
