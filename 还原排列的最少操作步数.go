package main

import "fmt"

func main() {
	fmt.Println(reinitializePermutation(6))
}

func reinitializePermutation(n int) int {
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = i
	}
	count := 0
	equal := func() bool {
		for i, num := range c {
			if i != num {
				return false
			}
		}
		return true
	}
	for count == 0 || !equal() {
		t := make([]int, n)
		for i := 0; i < n; i++ {
			if i&1 == 0 {
				t[i] = c[i>>1]
			} else {
				t[i] = c[n>>1+(i-1)>>1]
			}
		}
		c = t
		fmt.Println(c)
		count++
	}
	return count
}

func reinitializePermutationV2(n int) int {
	ans := 0
	for i := 1; ; {
		ans++
		if i&1 == 1 {
			i = n>>1 + i>>1
		} else if i&1 == 0 {
			i = i >> 1
		}
		if i == 1 {
			return ans
		}
	}
}
