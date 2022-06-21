package main

func main() {

}

func checkStraightLine(coordinates [][]int) bool {
	for i := 1; i < len(coordinates)-1; i++ {
		p1 := coordinates[i-1]
		p2 := coordinates[i]
		p3 := coordinates[i+1]
		if (p2[1]-p1[1])*(p3[0]-p2[0]) != (p3[1]-p2[1])*(p2[0]-p1[0]) {
			return false
		}
	}
	return true
}
