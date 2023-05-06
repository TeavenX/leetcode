package main

func main() {

}

func numPairsDivisibleBy60(time []int) int {
	ans := 0
	cache := [60]int{}
	for _, t := range time {
		ans += cache[(60-t%60)%60]
		cache[t%60]++
	}
	return ans
}
