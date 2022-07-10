package main

func main() {

}

func characterReplacement(s string, k int) int {
	left, right := 0, 0
	maxLen := 0
	cache := make(map[byte]int)
	diffLessThanK := func() bool {
		maxC := 0
		diffC := 0
		for _, v := range cache {
			if v > maxC {
				diffC += maxC
				maxC = v
			} else {
				diffC += v
			}
		}
		return diffC <= k
	}
	for right < len(s) {
		cache[s[right]]++
		right++
		if diffLessThanK() {
			l := right - left
			if l > maxLen {
				maxLen = l
			}
			continue
		}
		cache[s[left]]--
		left++
	}
	return maxLen
}
