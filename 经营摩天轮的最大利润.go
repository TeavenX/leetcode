package main

func main() {

}

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	ans := -1
	maxIncome := 0
	total := 0
	wait := 0
	i := 0
	for wait > 0 || i < len(customers) {
		if i < len(customers) {
			wait += customers[i]
		}
		num := min(wait, 4)
		wait -= num
		total += boardingCost*num - runningCost
		i++
		if total > maxIncome {
			maxIncome = total
			ans = i
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
