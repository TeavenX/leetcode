package main

func main() {

}

func captureForts(forts []int) int {
	ans := 0
	n := len(forts)
	for i := 0; i < n; {
		j := i + 1
		if forts[i] != 0 {
			for ; j < n && forts[j] == 0; j++ {
			}
			if j < n && forts[i]+forts[j] == 0 {
				ans = max(ans, j-i-1)
			}
		}
		i = j
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
