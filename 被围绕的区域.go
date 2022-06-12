package main

import "fmt"

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	//board = [][]byte{{'O'}}
	solve(board)
	fmt.Println(board)
}

var steps = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	maxLen := m
	if n > m {
		maxLen = n
	}
	queue := make([][]int, 0)
	for i := 0; i < maxLen; i++ {
		if i < n {
			if board[i][0] == 'O' {
				queue = append(queue, []int{i, 0})
				board[i][0] = 'H'
			}
			if board[i][m-1] == 'O' {
				queue = append(queue, []int{i, m - 1})
				board[i][m-1] = 'H'
			}
		}
		if i < m {
			if board[0][i] == 'O' {
				queue = append(queue, []int{0, i})
				board[0][i] = 'H'
			}
			if board[n-1][i] == 'O' {
				queue = append(queue, []int{n - 1, i})
				board[n-1][i] = 'H'
			}
		}
	}
	for len(queue) > 0 {
		level := len(queue)
		for idx := 0; idx < level; idx++ {
			i, j := queue[idx][0], queue[idx][1]
			for _, step := range steps {
				newi, newj := i+step[0], j+step[1]
				if 0 < newi && newi < n-1 && 0 < newj && newj < m-1 && board[newi][newj] == 'O' {
					board[newi][newj] = 'H'
					queue = append(queue, []int{newi, newj})
				}
			}
		}
		queue = queue[level:]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'H' {
				board[i][j] = 'O'
			}
		}
	}
}
