package main

func main() {

}

func maxRepOpt1(text string) int {
	cache := [26]int{}
	for _, b := range text {
		cache[b-'a']++
	}
	n := len(text)
	ans := 0
	for left := 0; left < n; {
		count := 1
		ct := cache[text[left]-'a']
		flag := true
		right := left + 1
		diff := right
		for ; right < n && count <= ct; right++ {
			if text[right] != text[left] {
				if ct-count > 0 && flag {
					diff = right
					ct--
					flag = false
				} else {
					break
				}
			}
			count++
		}
		if flag && ct-count > 0 {
			count++
		}
		ans = max(ans, count)
		left = diff
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
