package main

import "fmt"

func main() {
	s1 := "ta"
	s2 := "testabc"
	s1 = "hello"
	s2 = "ooolleoooleh"
	fmt.Println(checkInclusionV3(s1, s2))

}

func checkInclusion(s1 string, s2 string) bool {
	ctn1, ctn2 := [26]int{}, [26]int{}
	lenS1 := len(s1)
	for i := 0; i < lenS1; i++ {
		ctn1[s1[i]-'a']++
	}
	for i, str := range s2 {
		ctn2[str-'a']++
		for j := i + 1; j < i+lenS1 && j < len(s2); j++ {
			ctn2[s2[j]-'a']++
		}
		if ctn1 == ctn2 {
			return true
		}
		for j := i + 1; j < i+lenS1 && j < len(s2); j++ {
			ctn2[s2[j]-'a']--
		}
		ctn2[str-'a']--
	}
	return false
}

func checkInclusionV2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	ctn1, ctn2 := [26]int{}, [26]int{}
	lenS1 := len(s1)
	for i := 0; i < lenS1; i++ {
		ctn1[s1[i]-'a']++
	}
	for i, str := range s2 {
		ctn2[str-'a']++
		for j := i + 1; j < i+lenS1 && j < len(s2); j++ {
			ctn2[s2[j]-'a']++
		}
		if ctn1 == ctn2 {
			return true
		}
		ctn2 = [26]int{}
	}
	return false
}

func checkInclusionV3(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			// 说明当前长子串滑动窗口的字符个数超过了短子串个数
			// 将当前长子串滑动窗口左边的计数-1
			cnt[s2[left]-'a']--
			// 将滑动窗口的左指针向右滑动
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}
