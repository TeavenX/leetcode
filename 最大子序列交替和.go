package main

func main() {

}

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	f := make([][2]int, n+1)
	f[0][1] = -1e5
	for i, num := range nums {
		f[i+1][0] = max(f[i][0], f[i][1]-num)
		f[i+1][1] = max(f[i][1], f[i][0]+num)
	}
	return int64(max(f[n][0], f[n][1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
