package main

func main() {

}

const mod int = 1e9 + 7

func numDifferentIntegers(word string) int {
	cache := make(map[int]struct{})
	n := len(word)
	for i := 0; i < n; i++ {
		temp := 0
		flag := false
		for i < n && '0' <= word[i] && word[i] <= '9' {
			temp += int(word[i] - '0')
			temp = (temp * 10) % mod
			i++
			flag = true
		}
		if flag {
			cache[temp] = struct{}{}
			i--
		}
	}
	return len(cache)
}
