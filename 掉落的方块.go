package main

func main() {

}

func fallingSquares(positions [][]int) []int {
	n := len(positions)
	if n == 1 {
		return []int{positions[0][1]}
	}
	result := make([]int, n)
	for i, cube := range positions {
		left, right := cube[0], cube[0]+cube[1]-1
		result[i] = cube[1]
		for j := 0; j < i; j++ {
			preLeft, preRight := positions[j][0], positions[j][0]+positions[j][1]-1
			if right >= preLeft && preRight >= left {
				result[i] = max(result[i], result[j]+cube[1])
			}
		}
	}
	for i := 1; i < n; i++ {
		result[i] = max(result[i], result[i-1])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
