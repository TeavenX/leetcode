package main

func main() {

}

func minCount(coins []int) int {
	ans := 0
	for _, c := range coins {
		ans += (c + 1) >> 1
	}
	return ans
}
