package main

func main() {

}

func digitCount(num string) bool {
	cache := [10]int{}
	for _, b := range num {
		cache[b-'0']++
	}
	for i, b := range num {
		if cache[i] != int(b-'0') {
			return false
		}
	}
	return true
}
