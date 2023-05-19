package main

func main() {

}

func hammingDistance(x int, y int) int {
	r := x ^ y
	ans := 0
	for r > 0 {
		ans += r & 1
		r >>= 1
	}
	return ans
}

func hammingDistanceV2(x int, y int) int {
	r := x ^ y
	ans := 0
	for r > 0 {
		ans++
		r &= r - 1
	}
	return ans
}
