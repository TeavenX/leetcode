package main

func main() {

}

func canChoose(groups [][]int, nums []int) bool {
next:
	for _, group := range groups {
		for len(nums) >= len(group) {
			if sliceEqual(nums[:len(group)], group) {
				nums = nums[len(group):]
				continue next
			}
			nums = nums[1:]
		}
		return false
	}
	return true
}

func sliceEqual(a, b []int) bool {
	for i, num := range a {
		if num != b[i] {
			return false
		}
	}
	return true
}
