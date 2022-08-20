package main

func main() {

}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	result := 0
	for i, start := range startTime {
		if start <= queryTime && endTime[i] >= queryTime {
			result++
		}
	}
	return result
}
