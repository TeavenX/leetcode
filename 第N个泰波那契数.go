package main

func main() {

}

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	p, q, k, r := 0, 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = k
		k = r
		r = p + q + k
	}
	return r
}
