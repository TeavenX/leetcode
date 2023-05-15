package main

func main() {

}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	cache := make(map[[5]uint]int)
	ans := 0
	for _, row := range matrix {
		t := [5]uint{}
		for i, num := range row {
			t[i/64] |= uint((num ^ row[0]) << (i % 64))
		}
		cache[t]++
	}
	for _, v := range cache {
		ans = max(ans, v)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
