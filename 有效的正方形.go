package main

func main() {
	//p1 := []int{0, 0}
	//p2 := []int{-1, 0}
	//p3 := []int{1, 0}
	//p4 := []int{0, 1}

	//p1 := []int{1, 0}
	//p2 := []int{-1, 0}
	//p3 := []int{0, 1}
	//p4 := []int{0, -1}

	//p1 := []int{0, 0}
	//p2 := []int{1, 1}
	//p3 := []int{1, 0}
	//p4 := []int{0, 1}

	p1 := []int{1, 1}
	p2 := []int{5, 3}
	p3 := []int{3, 5}
	p4 := []int{7, 7}
	validSquare(p1, p2, p3, p4)
}

func check(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if pow(p3[0]-p1[0])+pow(p1[1]-p3[1]) == pow(p1[0]-p4[0])+pow(p1[1]-p4[1]) && pow(p1[0]-p4[0])+pow(p1[1]-p4[1]) == pow(p2[0]-p3[0])+pow(p2[1]-p3[1]) {
		x1 := p2[0] - p1[0]
		y1 := p2[1] - p1[1]
		x2 := p4[0] - p3[0]
		y2 := p4[1] - p3[1]
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			return false
		}
		if x1*x2 == -y1*y2 && pow(p1[0]-p2[0])+pow(p1[1]-p2[1]) == pow(p3[0]-p4[0])+pow(p3[1]-p4[1]) {
			return true
		}
	}
	return false
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if check(p1, p2, p3, p4) ||
		check(p1, p3, p2, p4) ||
		check(p1, p4, p2, p3) {
		return true
	}
	return false
}

func pow(x int) int {
	return x * x
}
