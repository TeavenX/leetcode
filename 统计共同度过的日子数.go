package main

import (
	"strconv"
	"strings"
)

func main() {

}

var days = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	return max(min(count(leaveAlice), count(leaveBob))-max(count(arriveAlice), count(arriveBob))+1, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func count(s string) int {
	m, d := split(s)
	return days[m-1] - days[0] + d
}

func split(s string) (int, int) {
	t := strings.Split(s, "-")
	return stringToInt(t[0]), stringToInt(t[1])
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
