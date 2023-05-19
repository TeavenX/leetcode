package main

func main() {

}

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&1 == 1 {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = ans[i>>1]
		}
	}
	return ans
}
