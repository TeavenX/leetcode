package main

func main() {

}

const (
	lowLetter = iota
	upLetter
	number
	special
)

func strongPasswordCheckerII(password string) bool {
	n := len(password)
	if n < 8 {
		return false
	}
	ans := 0
	for i := 0; i < n; i++ {
		b := password[i]
		if i > 0 && b == password[i-1] {
			return false
		}
		if b >= '0' && b <= '9' {
			ans |= 1 << number
		} else if b >= 'a' && b <= 'z' {
			ans |= 1 << lowLetter
		} else if b >= 'A' && b <= 'Z' {
			ans |= 1 << upLetter
		} else {
			ans |= 1 << special
		}
	}
	return ans == 15
}
