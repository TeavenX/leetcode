package main

func main() {
	maxLengthBetweenEqualCharacters("scayofdzca")
}

func maxLengthBetweenEqualCharacters(s string) int {
	l := 0
	cache := make(map[rune]int)
	for i, str := range s {
		if idx, exist := cache[str]; exist {
			if i-idx > l {
				l = i - idx
			}
		} else {
			cache[str] = i
		}
	}
	return l - 1
}
