package main

func main() {

}

func minStartValue(nums []int) int {
	sum := 1
	max := 1
	for _, num := range nums {
		sum -= num
		if sum > max {
			max = sum
		}
	}
	return max
}
