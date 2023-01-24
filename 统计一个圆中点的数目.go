package main

func main() {

}

func countPoints(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i, query := range queries {
		x, y, r := query[0], query[1], query[2]
		for _, point := range points {
			px, py := point[0], point[1]
			t := (x-px)*(x-px) + (y-py)*(y-py)
			if t <= r*r {
				ans[i]++
			}
		}
	}
	return ans
}
