package main

func main() {

}

func countTriplets(nums []int) int {
	ans := 0
	cache := make([]int, 1<<16) // 344ms, 6.8MB
	// cache := make(map[int]int) // 800ms, 7.6MB
	for _, x := range nums {
		for _, y := range nums {
			cache[x&y]++
		}
	}
	for _, z := range nums {
		for a, num := range cache {
			if num > 0 && z&a == 0 {
				ans += num
			}
		}
	}
	return ans
}
