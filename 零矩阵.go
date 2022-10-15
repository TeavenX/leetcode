package main

func main() {

}

func setZeroes(matrix [][]int) {
	r, c := make([]int, len(matrix[0])), make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				r[j] = 1
				c[i] = 1
			}
		}
	}
	for ci, row := range matrix {
		if c[ci] == 1 {
			for i := range row {
				row[i] = 0
			}
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if r[i] == 1 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
}
