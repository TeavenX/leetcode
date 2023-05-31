package main

import "sort"

func main() {

}

func maxDistance(position []int, m int) int {
	n := len(position)
	sort.Ints(position)
	if m == 2 {
		return position[n-1] - position[0]
	}
	if position[n-1] == position[0] {
		return 0
	}
	return sort.Search((position[n-1]-position[0])/(m-1)+1, func(d int) (res bool) {
		d++
		count := 1
		pre := position[0]
		for _, p := range position {
			if p >= pre+d {
				pre = p
				count++
			}
		}
		return count < m
	})
}
