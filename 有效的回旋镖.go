package main

func main() {

}

func isBoomerang(points [][]int) bool {
	k1 := (points[2][0] - points[0][0]) * (points[1][1] - points[0][1])
	k2 := (points[1][0] - points[0][0]) * (points[2][1] - points[0][1])
	return k1 != k2
}
