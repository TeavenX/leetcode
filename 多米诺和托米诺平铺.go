package main

func main() {

}

var mod int = 1e9 + 7

// https://pic.leetcode.cn/1668157188-nBzesC-790-5.png
// https://leetcode.cn/problems/domino-and-tromino-tiling/solution/by-endlesscheng-umpp/

func numTilings(n int) int {
	if n < 3 {
		return n
	}
	// f(n) = 2 * f(n-1) + f(n-3)
	a, b, c := 1, 1, 2
	for i := 3; i <= n; i++ {
		a, b, c = b, c, (2*c+a)%mod
	}
	return c
}
