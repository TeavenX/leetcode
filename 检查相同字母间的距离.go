package main

func main() {

}

func checkDistances(s string, distance []int) bool {
	cache := [26]int{}
	for i := range cache {
		cache[i] = -1
	}
	for i, b := range s {
		if cache[b-'a'] == -1 {
			cache[b-'a'] = i
		}
	}
	for i := 0; i < 26; i++ {
		if cache[i] != -1 {
			if cache[i]+distance[i]+1 >= len(s) || s[cache[i]] != s[cache[i]+distance[i]+1] {
				return false
			}
		}
	}
	return true
}

func checkDistancesV2(s string, distance []int) bool {
	cache := [26]int{}
	for i, b := range s {
		d := cache[b-'a']
		if d > 0 && i-d != distance[b-'a'] {
			return false
		}
		cache[b-'a'] = i + 1
	}
	return true
}
