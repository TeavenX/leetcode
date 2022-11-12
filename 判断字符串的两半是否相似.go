package main

func main() {

}

func halvesAreAlike(s string) bool {
	dict := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	c := make(map[byte]int)
	for _, b := range dict {
		c[b] = 1
	}
	n := len(s)
	count := 0
	for i := 0; i < n>>1; i++ {
		count += c[s[i]]
	}
	for i := n >> 1; i < n; i++ {
		count -= c[s[i]]
	}
	return count == 0
}
