package main

func main() {

}

func repeatedNTimes(nums []int) int {
	cache := make(map[int]bool)
	for _, num := range nums {
		if exist := cache[num]; exist {
			return num
		}
		cache[num] = true
	}
	return -1
}

func repeatedNTimesV2(nums []int) int {
	for idx, num := range nums[:len(nums)-1] {
		for i := 1; i <= 3 && idx+i < len(nums); i++ {
			if nums[idx+i] == num {
				return num
			}
		}
	}
	return -1
}
