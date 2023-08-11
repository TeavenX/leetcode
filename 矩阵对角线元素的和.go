package main

func main() {

}

func diagonalSum(mat [][]int) int {
	n := len(mat)
	ans := 0
	for i, j := 0, 0; i < n && j < n; i, j = i+1, j+1 {
		if i == n-1-i {
			ans += mat[i][j]
			continue
		}
		ans += mat[i][j] + mat[n-1-i][j]
	}
	return ans
}
