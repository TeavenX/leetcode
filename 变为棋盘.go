package main

import "fmt"

func main() {
	//board := [][]int{{1, 1, 0}, {0, 0, 1}, {0, 0, 1}}
	board := [][]int{{1, 0, 1, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 1, 0, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}}
	fmt.Println(movesToChessboard(board))
}

func movesToChessboardError(board [][]int) int {
	countR := 0
	countC := 0
	n := len(board)
	init := 0
	diffC, diffR := 0, 0
	line1 := 0
	for i := 0; i < n; i++ {
		line1 <<= 1
		line1 |= board[0][i]
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			linen := 0
			for j := 0; j < n; j++ {
				linen <<= 1
				linen |= board[i][j]
			}
			if linen != line1 && linen&line1 != 0 {
				return -1
			}
		}
		countC += board[0][i]
		countR += board[i][0]
		if init != board[0][i] {
			diffC++
		}
		if init != board[i][0] {
			diffR++
		}
		init = abs(init - 1)
	}
	if abs(n-2*countR) > 1 || abs(n-2*countC) > 1 {
		return -1
	}
	fmt.Println(diffR, diffC, n)
	if diffC+diffR < n {
		return (diffR + diffC + 1) >> 1
	} else {
		return (2*n - diffR - diffC + 1) >> 1
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func movesToChessboard(board [][]int) int {
	var (
		n       = len(board)
		rowSum  = 0
		colSum  = 0
		rowDiff = 0
		colDiff = 0
	)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (board[0][0])^(board[i][0])^(board[0][j])^(board[i][j]) != 0 {
				return -1
			}
		}
	}
	for i := 0; i < n; i++ {
		rowSum += board[0][i]
		colSum += board[i][0]
		if board[i][0] == i&1 {
			rowDiff++
		}
		if board[0][i] == i&1 {
			colDiff++
		}
	}
	if n/2 > rowSum || rowSum > (n+1)/2 {
		return -1
	}
	if n/2 > colSum || colSum > (n+1)/2 {
		return -1
	}
	if n&1 == 1 {
		if rowDiff&1 == 1 {
			rowDiff = n - rowDiff
		}
		if colDiff&1 == 1 {
			colDiff = n - colDiff
		}
	} else {
		rowDiff = min(n-rowDiff, rowDiff)
		colDiff = min(n-colDiff, colDiff)
	}
	return (rowDiff + colDiff) / 2

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
