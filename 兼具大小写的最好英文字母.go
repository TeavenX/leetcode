package main

func main() {

}

func greatestLetter(s string) string {
	cache := make(map[rune]struct{})
	for _, b := range s {
		cache[b] = struct{}{}
	}
	for i := 'Z'; i >= 'A'; i-- {
		if _, ok := cache[rune(i)]; ok {
			if _, ok = cache[i+32]; ok {
				return string(i)
			}
		}
	}
	return ""
}
