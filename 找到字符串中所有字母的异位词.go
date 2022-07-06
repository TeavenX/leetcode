package main

import "fmt"

func main() {
	fmt.Println(findAnagrams20220707("cbaebabacd", "abc"))
	//fmt.Println(findAnagrams("abab", "ab"))
	//fmt.Println(findAnagrams("baa", "aa"))
}

func findAnagramsTIMEOUT(s string, p string) []int {
	var backtracing func()
	cache := make(map[string]bool)
	cacheIdx := make(map[int]bool)
	path := ""
	k := len(p)
	backtracing = func() {
		if len(path) == k {
			cache[path] = true
		}
		for i := 0; i < k; i++ {

			if used := cacheIdx[i]; !used {
				path += string(p[i])
				cacheIdx[i] = true
				backtracing()
				cacheIdx[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtracing()
	result := make([]int, 0)
	for i := 0; i <= len(s)-k; i++ {
		str := s[i : i+k]
		if exists := cache[str]; exists {
			result = append(result, i)
		}
	}
	return result
}

func findAnagrams(s string, p string) []int {
	window, need := make(map[byte]int), make(map[byte]int)
	left, right, valid := 0, 0, 0
	result := make([]int, 0)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	for right < len(s) {
		str := s[right]
		right++
		if _, exists := need[str]; exists {
			window[str]++
			if window[str] == need[str] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
			}
			strL := s[left]
			left++
			if _, exists := need[strL]; exists {
				if window[strL] == need[strL] {
					valid--
				}
				window[strL]--
			}
		}
	}
	return result
}

func findAnagrams20220707(s string, p string) []int {
	size := len(p)
	if len(s) < size {
		return []int{}
	}
	cache, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < size; i++ {
		cache[p[i]]++
	}
	result := make([]int, 0)
	left, right, valid := 0, 0, 0
	need := len(cache)
	for right < len(s) {
		strR := s[right]
		right++
		if _, exists := cache[strR]; exists {
			window[strR]++
			if window[strR] == cache[strR] {
				valid++
			}
		}
		if right-left == size {
			if valid == need {
				result = append(result, left)
			}
			strL := s[left]
			left++
			if _, exists := cache[strL]; exists {
				if window[strL] == cache[strL] {
					valid--
				}
				window[strL]--
			}
		}
	}
	return result
}
