package main

func main() {

}

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		b := s[left]
		for left < right && s[left] == b {
			left++
		}
		for left <= right && s[right] == b {
			right--
		}
	}
	return right - left + 1
}
