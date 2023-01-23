package main

func main() {

}

func calculateTax(brackets [][]int, income int) float64 {
	ans := 0
	for i := 0; i < len(brackets); i++ {
		u, p := brackets[i][0], brackets[i][1]
		var preU int
		if i > 0 {
			preU = brackets[i-1][0]
		}
		if u < income {
			ans += (u - preU) * p
		} else {
			ans += (income - preU) * p
			break
		}
	}
	return float64(ans) / 100
}

func calculateTaxV2(brackets [][]int, income int) float64 {
	ans, pre := 0, 0
	for _, b := range brackets {
		t := (min(b[0], income) - pre) * b[1]
		if t < 0 {
			break
		}
		ans += t
		pre = b[0]
	}
	return float64(ans) / 100
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
