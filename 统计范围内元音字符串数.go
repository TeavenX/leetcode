package main

func main() {

}

var mp = map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}

func vowelStrings(words []string, queries [][]int) []int {
	cache := make([]int, len(words)+1)
	for i, word := range words {
		cache[i+1] = cache[i] + (mp[word[0]]+mp[word[len(word)-1]])>>1
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = cache[q[1]+1] - cache[q[0]]
	}
	return ans
}
