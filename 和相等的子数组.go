package main

func main() {

}

func findSubarrays(nums []int) bool {
	cache := make(map[int]bool)
	for i := 1; i < len(nums); i++ {
		n := nums[i] + nums[i-1]
		if cache[n] {
			return true
		}
		cache[n] = true
	}
	return false
}
