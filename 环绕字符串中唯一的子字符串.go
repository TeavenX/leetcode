package main

import "fmt"

func main() {
	//fmt.Println(findSubstringInWraproundString("zab"))
	//fmt.Println(findSubstringInWraproundString("cac"))
	//fmt.Println(findSubstringInWraproundString("zac"))
	fmt.Println(findSubstringInWraproundString("abaab"))
	fmt.Println(findSubstringInWraproundString("abcabc"))
}

func findSubstringInWraproundStringError(p string) int {
	n := len(p)
	if n == 1 {
		return 1
	}
	result := 0
	temp := 1
	cache := make(map[string]int)
	count := func(str string) {
		if _, ok := cache[str]; !ok {
			cache[str] = ((1 + temp) * temp) >> 1
			if len(str) > 1 {
				for _, s := range str {
					cache[string(s)] = 0
				}
			}
		}
	}
	for i := 1; i < n; i++ {
		if p[i]-p[i-1] == 1 || (p[i] == 'a' && p[i-1] == 'z') {
			temp++
		} else {
			str := p[i-temp : i]
			count(str)
			temp = 1
		}
	}
	count(p[n-temp : n])
	for _, v := range cache {
		result += v
	}
	return result
}

func findSubstringInWraproundString(p string) int {
	n := len(p)
	if n == 1 {
		return 1
	}
	dp := [26]int{}
	temp := 0
	for i, str := range p {
		if i > 0 && (byte(str)-p[i-1]+26)%26 == 1 {
			temp++
		} else {
			temp = 1
		}
		dp[str-'a'] = max(dp[str-'a'], temp)
	}
	result := 0
	for _, v := range dp {
		result += v
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
