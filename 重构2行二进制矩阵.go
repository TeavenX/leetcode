package main

func main() {

}

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	ans := make([][]int, 2)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i, s := range colsum {
		if s == 2 {
			ans[0][i] = 1
			ans[1][i] = 1
			upper--
			lower--
		}
		if upper < 0 || lower < 0 {
			return [][]int{}
		}
	}
	for i, s := range colsum {
		if s == 1 {
			if upper > 0 {
				ans[0][i] = 1
				upper--
			} else if lower > 0 {
				ans[1][i] = 1
				lower--
			} else {
				return [][]int{}
			}
		}
	}
	if upper > 0 || lower > 0 {
		return [][]int{}
	}
	return ans
}
