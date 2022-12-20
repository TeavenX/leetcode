package main

func main() {

}

func maximumScore(a int, b int, c int) int {
	// 当有其中一堆石子的数量超过了另外两堆的总和，
	// 那么另外两堆总是能被取完
	// 否则一共可以被取三堆总数的一半次
	if a+b <= c {
		return a + b
	} else if a+c <= b {
		return a + c
	} else if b+c <= a {
		return b + c
	}
	return (a + b + c) / 2
}
