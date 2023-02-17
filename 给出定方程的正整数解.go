package main

func main() {

}

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	ans := make([][]int, 0)
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 1000; j++ {
			if customFunction(i, j) == z {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func findSolutionError(customFunction func(int, int) int, z int) [][]int {
	ans := make([][]int, 0)
	right := 10001000
	left := 10001
	for left < right {
		mid := (left + right) >> 1
		x, y := mid%1e4, mid/1e4
		r := customFunction(x, y)
		if r == z {
			ans = append(ans, []int{x, y})
			left++
		} else if r > z {
			right = mid
		} else {
			left = mid
		}
	}
	return ans
}

func findSolution(customFunction func(int, int) int, z int) [][]int {
	ans := make([][]int, 0)
	x, y := 1, 1000
	for x <= 1000 && y > 0 {
		r := customFunction(x, y)
		if r == z {
			ans = append(ans, []int{x, y})
			x++
			y--
		} else if r < z {
			x++
		} else {
			y--
		}
	}
	return ans
}
