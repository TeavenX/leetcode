package main

func main() {

}

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	var cache int32 = 0
	var cmp int32 = 1<<26 - 1
	for _, s := range sentence {
		cache |= 1 << (s - 'a')
	}
	return cmp == cache
}
