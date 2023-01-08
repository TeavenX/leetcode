package main

func main() {

}

func prefixCount(words []string, pref string) int {
	count := 0
	n := len(pref)
	for _, word := range words {
		if n <= len(word) && word[:n] == pref {
			count++
		}
	}
	return count
}
