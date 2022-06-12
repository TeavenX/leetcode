package main

func main() {

}

func diStringMatch(s string) []int {
	n := len(s)
	result := make([]int, n+1)
	left, right, idx := 0, n, 0
	for left < right {
		if s[idx] == 'I' {
			result[idx] = left
			left++
		} else {
			result[idx] = right
			right--
		}
		idx++
	}
	result[idx] = left
	return result
}
