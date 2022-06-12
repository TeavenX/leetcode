package main

func main() {

}

func fib(n int) int {
	if n < 2 {
		return n
	}
	cache := map[int]int{0: 0, 1: 1}
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}

func fibV2(n int) int {
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

func fib20220504(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 1, 1
	for i := 2; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
