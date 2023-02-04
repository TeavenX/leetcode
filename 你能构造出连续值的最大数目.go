package main

func main() {

}

func getMaximumConsecutive(coins []int) int {
	ans := 0
	sort.Ints(coins)
	for _, coin := range coins {
		if coin > ans+1 {
			break
		}
		ans += coin
	}
	return ans + 1
}
