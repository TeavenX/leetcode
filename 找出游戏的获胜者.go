package main

import "fmt"

func main() {
	fmt.Println(findTheWinner(6, 5))
}

func findTheWinner(n int, k int) int {
	pos := 0
	for i := 2; i <= n; i++ {
		pos = (pos + k) % i
	}
	return pos + 1
}
