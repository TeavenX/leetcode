package main

func main() {

}

func maxScoreWords(words []string, letters []byte, score []int) int {
	cache := [26]int{}
	for _, let := range letters {
		cache[let-'a']++
	}
	ans := 0
	var traceback func(idx int, s int)
	traceback = func(idx, s int) {
		if idx == len(words) {
			ans = max(ans, s)
			return
		}
		tmp := 0
		i := len(words[idx])
		for idx, let := range words[idx] {
			if cache[let-'a'] == 0 {
				tmp = 0
				i = idx
				break
			}
			cache[let-'a']--
			tmp += score[let-'a']
		}
		for i := idx + 1; i <= len(words); i++ {
			traceback(i, s+tmp)
		}
		for _, let := range words[idx][:i] {
			cache[let-'a']++
		}
	}
	for i := range words {
		traceback(i, 0)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
