package main

func main() {

}

func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	n, m := len(first), len(second)
	dist := n - m
	if dist <= -2 || dist >= 2 {
		return false
	}
	if n > m {
		first, second = second, first
		n, m = m, n
	}
	dist = 0
	i, j := 0, 0
	for i < n && j < m {
		if first[i] == second[j] {
			i++
			j++
		} else {
			if n == m {
				i++
				j++
				dist++
			} else {
				dist++
				if dist > 1 {
					return false
				}
				j++
			}
		}
	}
	return dist <= 1
}
