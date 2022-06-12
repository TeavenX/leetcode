package main

func main() {

}

func findDuplicates(nums []int) []int {
	result := make([]int, 0)
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] < 0 {
			result = append(result, num)
		} else {
			nums[num-1] = -nums[num-1]
		}
	}
	return result
}
