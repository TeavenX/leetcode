package main

import "sort"

func main() {

}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	result := 0
	for i, seat := range seats {
		result += abs(seat - students[i])
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
