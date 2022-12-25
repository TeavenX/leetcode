package main

func main() {

}

const mod int = 1e9 + 7

func countHomogenous(s string) int {
	count := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if s[right] != s[left] {
			c := right - left
			count = (count + (1+c)*c>>1) % mod
			left = right
		}
	}
	c := len(s) - left
	count = (count + (1+c)*c>>1) % mod
	return count
}
