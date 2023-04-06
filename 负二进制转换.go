package main

import "fmt"

func main() {
	n := 2
	k := 1
	ans := make([]byte, 0)
	for n != 0 {
		if n%2 != 0 {
			ans = append(ans, '1')
			n -= k
		} else {
			ans = append(ans, '0')
		}
		k *= -1
		n /= 2
	}
	fmt.Println(string(ans))
}

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	ans := make([]byte, 0)
	k := 1
	for n != 0 {
		if n&1 == 1 {
			ans = append(ans, '1')
			n -= k
		} else {
			ans = append(ans, '0')
		}
		k *= -1
		n /= 2
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
