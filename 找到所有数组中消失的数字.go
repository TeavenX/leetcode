package main

func main() {

}

func findDisappearedNumbers(nums []int) []int {
	for _, i := range nums {
		if i < 0 {
			i = -i
		}
		i--
		if nums[i] > 0 {
			nums[i] = -nums[i]
		}
	}
	ans := make([]int, 0)
	for i, num := range nums {
		if num > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
