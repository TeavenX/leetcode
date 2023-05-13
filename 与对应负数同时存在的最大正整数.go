package main

func main() {

}

func findMaxK(nums []int) int {
	cache := make(map[int]bool)
	ans := -1
	for _, num := range nums {
		if cache[-num] {
			ans = max(ans, abs(num))
		}
		cache[num] = true
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
