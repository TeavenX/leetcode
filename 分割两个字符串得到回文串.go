package main

func main() {

}

func checkPalindromeFormation(a string, b string) bool {
	return checkPalindrome(a, b) || checkPalindrome(b, a)
}

func checkPalindrome(a, b string) bool {
	left, right := 0, len(a)-1
	for left < right && a[left] == b[right] {
		left++
		right--
	}
	return check(a, left, right) || check(b, left, right)
}

func check(s string, left, right int) bool {
	for left < right && s[left] == s[right] {
		left++
		right--
	}
	return left >= right
}
