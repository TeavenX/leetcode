package main

func main() {

}

func minimumMoves(s string) int {
	n := len(s)
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == 'X' {
			count++
			i += 2
		}
	}
	return count
}
