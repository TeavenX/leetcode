package main

func main() {

}

func areAlmostEqual(s1 string, s2 string) bool {
	wm := make(map[rune]int)
	for _, str := range s1 {
		wm[str]++
	}
	for _, str := range s2 {
		wm[str]--
	}
	for _, v := range wm {
		if v != 0 {
			return false
		}
	}
	dist := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			dist++
		}
		if dist > 2 {
			return false
		}
	}
	return true
}

func areAlmostEqualV2(s1 string, s2 string) bool {
	wm := make(map[byte]int)
	dist := 0
	for i := 0; i < len(s1); i++ {
		wm[s1[i]]++
		wm[s2[i]]--
		if s1[i] != s2[i] {
			dist++
			if dist > 2 {
				return false
			}
		}
	}
	for _, v := range wm {
		if v != 0 {
			return false
		}
	}
	return true
}

func areAlmostEqualV3(s1 string, s2 string) bool {
	dist := 0
	c1, c2 := byte('#'), byte('#')
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if c1 == '#' {
				c1 = s1[i]
				c2 = s2[i]
				dist++
			} else {
				if c2 == s1[i] && c1 == s2[i] {
					dist--
					continue
				}
				return false
			}
		}
	}
	if dist != 0 {
		return false
	}
	return true
}

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	i, j := -1, -1
	for idx := 0; idx < len(s1); idx++ {
		if s2[idx] != s1[idx] {
			if i == -1 {
				i = idx
			} else if j == -1 {
				j = idx
			} else {
				return false
			}
		}
	}
	if j == -1 {
		return false
	}
	if s2[i] != s1[j] || s1[i] != s2[j] {
		return false
	}
	return true
}
