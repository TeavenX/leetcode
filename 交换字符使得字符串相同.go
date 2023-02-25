package main

func main() {

}

func minimumSwap(s1 string, s2 string) int {
	cx, cy := 0, 0
	for i, s := range s1 {
		if s != rune(s2[i]) {
			if s == 'x' {
				cx++
			} else {
				cy++
			}
		}
	}
	d := cx + cy
	if d&1 == 1 {
		return -1
	}
	if cx&1 == 0 {
		return d >> 1
	}
	return d>>1 + 1
}
