package main

func main() {

}

func numSpecial(mat [][]int) int {
	l, c := make([]int, 0), make([]int, 0)
	for i, line := range mat {
		s := 0
		for _, n := range line {
			s += n
		}
		if s == 1 {
			l = append(l, i)
		}
	}
	for i, _ := range mat[0] {
		s := 0
		for j := 0; j < len(mat); j++ {
			s += mat[j][i]
		}
		if s == 1 {
			c = append(c, i)
		}
	}
	result := 0
	for _, cn := range c {
		for _, ln := range l {
			result += mat[ln][cn]
		}
	}
	return result
}
