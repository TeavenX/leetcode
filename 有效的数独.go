package main

import "fmt"

func main() {
	board := make([][]byte, 9)
	board[0] = []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
	board[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	board[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	board[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	board[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	board[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	board[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	board[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	board[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	cacheC := make([]int, 9)
	cacheR := make([]int, 9)
	cacheBlock := make([][]int, 3)
	for i := 0; i < 9; i++ {
		n := 1 << 10
		cacheC[i] = n
		cacheR[i] = n
		if i < 3 {
			cacheBlock[i] = []int{n, n, n}
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				if ((cacheC[i]>>num)|(cacheR[j]>>num)|(cacheBlock[i/3][j/3]>>num))&1 == 1 {
					return false
				}
				n := 1 << num
				cacheC[i] |= n
				cacheR[j] |= n
				cacheBlock[i/3][j/3] |= n
			}
		}
	}
	return true
}

func isValidSudokuV2(board [][]byte) bool {
	cacheC := make([]int, 9)
	cacheR := make([]int, 9)
	cacheBlock := make([][]int, 3)
	for i := 0; i < 3; i++ {
		cacheBlock[i] = make([]int, 3)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				if ((cacheC[i]>>num)|(cacheR[j]>>num)|(cacheBlock[i/3][j/3]>>num))&1 == 1 {
					return false
				}
				n := 1 << num
				cacheC[i] |= n
				cacheR[j] |= n
				cacheBlock[i/3][j/3] |= n
			}
		}
	}
	return true
}
