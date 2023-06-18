package main

func main() {

}

func maxSumDivThree(nums []int) int {
	cache := [3]int{}
	for _, num := range nums {
		a := cache[0] + num
		b := cache[1] + num
		c := cache[2] + num
		cache[a%3] = max(cache[a%3], a)
		cache[b%3] = max(cache[b%3], b)
		cache[c%3] = max(cache[c%3], c)
	}
	return cache[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
