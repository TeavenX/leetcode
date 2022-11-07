package main

func main() {

}

func countConsistentStrings(allowed string, words []string) int {
	cache := [26]bool{}
	for _, s := range allowed {
		cache[s-'a'] = true
	}
	result := 0
	for _, word := range words {
		flag := true
		for _, s := range word {
			if !cache[s-'a'] {
				flag = false
				break
			}
		}
		if flag {
			result++
		}
	}
	return result
}
