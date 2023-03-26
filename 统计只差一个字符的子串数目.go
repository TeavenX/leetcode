package main

func main() {

}

func countSubstrings(s string, t string) int {
	ans := 0
	n, m := len(s), len(t)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			diff := 0
			for k := 0; i+k < n && j+k < m; k++ {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff > 1 {
					break
				}
				if diff == 1 {
					ans++
				}
			}
		}
	}
	return ans
}
