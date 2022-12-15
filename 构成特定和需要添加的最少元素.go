package main

func main() {

}

func minElements(nums []int, limit int, goal int) int {
	for _, num := range nums {
		goal -= num
	}
	if goal < 0 {
		goal = -goal
	}
	count := 0
	for goal > 0 {
		if goal >= limit {
			goal -= limit
		} else {
			goal = 0
		}
		count++
	}
	return count
}

func minElementsV2(nums []int, limit int, goal int) int {
	for _, num := range nums {
		goal -= num
	}
	if goal < 0 {
		goal = -goal
	}
	return (goal + limit - 1) / limit
}
