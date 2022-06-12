package main

func main() {

}

var steps = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	n := len(word)
	nb := len(board)
	mb := len(board[0])
	var traceback func(k, i, j int) bool
	traceback = func(k, i, j int) bool {
		if k == n-1 {
			return board[i][j] == word[k]
		}
		if board[i][j] == word[k] {
			temp := board[i][j]
			board[i][j] = ' '
			for _, step := range steps {
				newi, newj := i+step[0], j+step[1]
				if 0 <= newi && newi < nb && 0 <= newj && newj < mb {
					if traceback(k+1, newi, newj) {
						return true
					}
				}
			}
			board[i][j] = temp
		}
		return false
	}
	for i := 0; i < nb; i++ {
		for j := 0; j < mb; j++ {
			if traceback(0, i, j) {
				return true
			}
		}
	}
	return false
}
