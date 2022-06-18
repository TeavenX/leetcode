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
