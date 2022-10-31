package main

func main() {

}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	l1, l2 := 0, 0
	i1, i2 := 0, 0
	n1, n2 := len(word1), len(word2)
	r1, r2 := "", ""
	for i1 < n1 || i2 < n2 {
		if i1 < n1 {
			r1 += word1[i1]
			l1 += len(word1[i1])
			i1++
		}
		if i2 < n2 {
			r2 += word2[i2]
			l2 += len(word2[i2])
			i2++
		}
		l := min(l1, l2)
		if r1[:l] != r2[:l] {
			return false
		}
	}
	if l1 != l2 {
		return false
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
