package main

func main() {

}

func majorityElement(nums []int) int {
	ans := -1
	sum := 0
	for _, num := range nums {
		if ans == -1 {
			ans = num
		}
		if ans == num {
			sum++
		} else {
			sum--
		}
		if sum == 0 {
			ans = -1
		}
	}
	return ans
}
