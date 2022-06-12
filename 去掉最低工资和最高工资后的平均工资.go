package main

import "fmt"

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	fmt.Println(average(salary))
}

func average(salary []int) float64 {
	sum := 0
	low, high := salary[0], salary[1]
	if salary[0] > salary[1] {
		low, high = salary[1], salary[0]
	}
	for _, num := range salary[2:] {
		if num < low {
			sum += low
			low = num
		} else if num > high {
			sum += high
			high = num
		} else {
			sum += num
		}
	}
	return float64(sum) / float64(len(salary)-2)
}
