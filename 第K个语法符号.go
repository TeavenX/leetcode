package main

import "fmt"

func main() {
	fmt.Println(kthGrammar(4, 1))
}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	return k&1 ^ 1 ^ (kthGrammar(n-1, (k+1)>>1))
}
