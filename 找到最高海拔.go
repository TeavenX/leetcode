package main

func main() {

}

func largestAltitude(gain []int) int {
	max, n := 0, 0
	for _, h := range gain {
		n += h
		if n > max {
			max = n
		}
	}
	return max
}
