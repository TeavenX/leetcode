package main

func main() {

}

// func numMatchingSubseq(s string, words []string) int {
//     n := len(s)
//     count := 0
//     for _, word := range words {
//         i, j := 0, 0
//         for i < n && j < len(word) {
//             if word[j] == s[i] {
//                 j++
//             }
//             i++
//         }
//         if j == len(word) {
//             count++
//         }
//     }
//     return count
// }

func numMatchingSubseq(s string, words []string) int {
	n := len(s)
	count := 0
	cache := [26][]int{}
	for i, str := range s {
		cache[str-'a'] = append(cache[str-'a'], i)
	}
	for _, word := range words {
		i, j := 0, 0
		for i < n && j < len(word) {
			c := cache[word[j]-'a']
			if len(c) == 0 || c[len(c)-1] < i {
				break
			}
			if i == 0 {
				i = c[0]
			} else {
				l, r := 0, len(c)-1
				for l < r {
					mid := (l + r) >> 1
					if c[mid] <= i {
						l = mid + 1
					} else {
						r = mid
					}
				}
				if c[l] > i {
					i = c[l]
				} else {
					break
				}
			}
			j++
		}
		if j == len(word) {
			count++
		}
	}
	return count
}
