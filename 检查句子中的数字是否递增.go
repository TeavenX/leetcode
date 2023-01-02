package main

func main() {

}

func areNumbersAscending(s string) bool {
	num := 0
	t := 0
	for _, b := range s {
		if b == ' ' && t > 0 {
			if t <= num {
				return false
			}
			num = t
			t = 0
		} else if b >= '0' && b <= '9' {
			t = t*10 + int(b-'0')
		}
	}
	if t > 0 && t <= num {
		return false
	}
	return true
}
