package main

func main() {

}

func climbStairs(n int) int {
	n = n + 1
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func climbStairs20220705(n int) int {
	if n <= 3 {
		return n
	}
	p, q, r := 1, 2, 0
	for i := 2; i < n; i++ {
		r = p + q
		p = q
		q = r
	}
	return r
}
