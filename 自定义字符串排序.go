package main

func main() {

}

// func customSortString(order string, s string) string {
//     cache := [26]int{}
//     for i := range order {
//         cache[order[i]-'a'] = i+1
//     }
//     b := []byte(s)
//     sort.Slice(b, func(i, j int) bool {
//         return cache[b[i]-'a'] < cache[b[j]-'a']
//     })
//     return string(b)
// }

func customSortString(order string, s string) string {
	cache := [26]int{}
	for _, b := range s {
		cache[b-'a']++
	}
	result := make([]byte, 0, len(s))
	for _, b := range order {
		for i := 0; i < cache[b-'a']; i++ {
			result = append(result, byte(b))
		}
		cache[b-'a'] = 0
	}
	for i := 'a'; i <= 'z'; i++ {
		for j := 0; j < cache[i-'a']; j++ {
			result = append(result, byte(i))
		}
	}
	return string(result)
}
