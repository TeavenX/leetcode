package main

func main() {

}

func camelMatch(queries []string, pattern string) []bool {
	ans := make([]bool, len(queries))
next:
	for idx, q := range queries {
		j := 0
		for i := 0; i < len(q); i++ {
			for j < len(pattern) && pattern[j] == q[i] {
				j++
				i++
			}
			if i < len(q) && q[i] >= 'A' && q[i] <= 'Z' {
				ans[idx] = false
				continue next
			}
		}
		ans[idx] = j == len(pattern)
	}
	return ans
}
