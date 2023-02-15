package main

func main() {

}

func numberOfPairs(nums []int) []int {
	cache := [101]bool{}
	n := len(nums)
	for _, num := range nums {
		cache[num] = !cache[num]
		if !cache[num] {
			n -= 2
		}
	}
	return []int{(len(nums) - n) >> 1, n}
}
