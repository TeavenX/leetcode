package main

import (
	"fmt"
)

func main() {
	//fmt.Println('0', '9', 'A', 'Z', 'a', 'z')
	//fmt.Printf("%s", string('A'))
	//fmt.Println()
	//fmt.Println("test"[2])
	//fmt.Println("test"[:1])
	//fmt.Println(string(97))
	a := "a1b2"
	//a = "C"
	fmt.Println(letterCasePermutationV2(a))
}

func letterCasePermutation(s string) (result []string) {
	n := len(s)
	distance := 'a' - 'A'
	cache := make(map[string]bool)
	var traceback func(string)
	traceback = func(prefix string) {
		newN := len(prefix)
		for _, str := range s[newN:] {
			if str <= 57 {
				prefix += string(str)
			} else {
				traceback(prefix + string(str))
				if str <= 90 {
					prefix += string(str + distance)
					traceback(prefix)
				} else if str <= 122 {
					prefix += string(str - distance)
					traceback(prefix)
				}
			}
			if exist := cache[prefix]; !exist && len(prefix) == n {
				result = append(result, prefix)
				cache[prefix] = true
				return
			}
		}
	}
	result = append(result, s)
	cache[s] = true
	traceback("")
	return
}

func letterCasePermutationV2(s string) (result []string) {
	n := len(s)
	str := []byte(s)
	var traceback func(int)
	traceback = func(startIdx int) {
		result = append(result, string(str))
		for i := startIdx; i < n; i++ {
			if 'a' <= str[i] && str[i] <= 'z' {
				str[i] -= 32
				traceback(i + 1)
				str[i] += 32
			}
			if 'A' <= str[i] && str[i] <= 'Z' {
				str[i] += 32
				traceback(i + 1)
				str[i] -= 32
			}
		}
	}
	traceback(0)
	return
}
