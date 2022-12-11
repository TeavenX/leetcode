package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(beautySumV3("aabcb"))
}

func beautySum(s string) int {
	tmp := make([]string, 0)
	for right := range s {
		left := 0
		for right-left >= 2 {
			tmp = append(tmp, s[left:right+1])
			left++
		}
	}
	result := 0
	for _, str := range tmp {
		cache := [26]int{}
		for _, t := range str {
			cache[t-'a']++
		}
		max, min := 0, math.MaxInt
		for i := 0; i < 26; i++ {
			if c := cache[i]; c > 0 {
				if c > max {
					max = c
				}
				if c < min {
					min = c
				}
			}
		}
		result += max - min
	}
	return result
}

func beautySumV2(s string) int {
	n := len(s)
	cache := make([][26]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 0; j < 26; j++ {
			cache[i][j] = cache[i-1][j]
		}
		cache[i][s[i-1]-'a']++
	}
	result := 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			max, min := 0, math.MaxInt
			for k := 0; k < 26; k++ {
				if v := cache[j][k] - cache[i-1][k]; v > 0 {
					if v > max {
						max = v
					}
					if v < min {
						min = v
					}
				}
			}
			result += max - min
		}
	}
	return result
}

func beautySumV3(s string) int {
	n := len(s)
	result := 0
	for i := 0; i < n; i++ {
		cache := [26]int{}
		for j := i; j < n; j++ {
			cache[s[j]-'a']++
			max, min := 0, math.MaxInt
			for _, v := range cache {
				if v > 0 {
					if v > max {
						max = v
					}
					if v < min {
						min = v
					}
				}
			}
			result += max - min
		}
	}
	return result
}
