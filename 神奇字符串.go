package main

func main() {
	magicalString(10)
}

func magicalString(n int) int {
	if n <= 3 {
		return 1
	}
	count := 1
	b := []int{1, 2, 2}
	slow, fast := 2, 3
	for fast < n {
		t := gen(b[fast-1])
		for i := 0; i < b[slow] && fast < n; i++ {
			b = append(b, t)
			if t == 1 {
				count++
			}
			fast++
		}
		slow++
	}
	return count
}

func gen(a int) int {
	if a == 1 {
		return 2
	}
	return 1
}

func magicalString(n int) int {
	if n <= 3 {
		return 1
	}
	count := 1
	b := make([]int, n)
	b[0] = 1
	b[1] = 2
	b[2] = 2
	slow, fast := 2, 3
	for fast < n {
		t := b[fast-1] ^ 3
		for i := 0; i < b[slow] && fast < n; i++ {
			b[fast] = t
			if t == 1 {
				count++
			}
			fast++
		}
		slow++
	}
	return count
}
