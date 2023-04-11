package main

func main() {

}

func longestDecomposition(text string) int {
	n := len(text)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1
	for left < right {
		if text[0:left+1] == text[right:n] {
			return 2 + longestDecomposition(text[left+1:right])
		}
		left++
		right--
	}
	return 1
}
