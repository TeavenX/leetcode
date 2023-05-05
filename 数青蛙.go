package main

func main() {

}

var m = map[rune]int{'c': 0, 'r': 1, 'o': 2, 'a': 3, 'k': 4}

func minNumberOfFrogs(croakOfFrogs string) int {
	unfc, mx := 0, 0
	count := [5]int{}
	for _, c := range croakOfFrogs {
		count[m[c]]++
		if c == 'c' {
			unfc++
			mx = max(mx, unfc)
		} else if c == 'k' {
			unfc--
		}
		for i := 0; i < m[c]; i++ {
			if count[i] < count[m[c]] {
				return -1
			}
		}
	}
	c := count[0]
	for _, cnt := range count {
		if c != cnt {
			return -1
		}
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
