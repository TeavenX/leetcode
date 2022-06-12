package main

func main() {

}

func countOdds(low int, high int) int {
	result := 0
	for low <= high {
		result += low & 1
		low++
	}
	return result
}

func countOddsV2(low int, high int) int {
	result := (high - low + 1) >> 1
	if (high-low)&1 == 1 && low&1 == 1 {
		result++
	}
	return result
}

func countOddsV3(low int, high int) int {
	return (high+1)>>1 - low>>1
}

func countOddsV4(low int, high int) int {
	countOdd := func(x int) int {
		return (x + 1) >> 1
	}
	return countOdd(high) - countOdd(low-1)
}
