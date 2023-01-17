package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(areSentencesSimilar("of", "A lot of words"))
	// fmt.Println(areSentencesSimilar("c h p Ny", "c BDQ r h p Ny"))
	// fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	// fmt.Println(areSentencesSimilar("hsYZKp Cn eE", "hsYZKp eE"))
	fmt.Println(areSentencesSimilar("a", "a aa a"))
}

func areSentencesSimilarErr(sentence1 string, sentence2 string) bool {
	if sentence1 == sentence2 {
		return true
	}
	if len(sentence1) > len(sentence2) {
		return areSentencesSimilar(sentence2, sentence1)
	}
	if strings.HasPrefix(sentence2, sentence1) || strings.HasSuffix(sentence2, sentence1) {
		return true
	}
	s1, s2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	if len(s1) < 2 {
		return false
	}
	l1, r1 := 0, len(s1)-1
	l2, r2 := 0, len(s2)-1
	for l1 <= r1 && l2 < r2 {
		if s1[l1] == s2[l2] {
			l1++
		}
		l2++
		if s1[r1] == s2[r2] {
			r1--
		}
		r2--
	}
	return l1-r1 == 1 && l2 != r2
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if sentence1 == sentence2 {
		return true
	}
	if len(sentence1) > len(sentence2) {
		return areSentencesSimilar(sentence2, sentence1)
	}
	s1, s2 := strings.Fields(sentence1), strings.Fields(sentence2)
	l, r := 0, 0
	l1, l2 := len(s1), len(s2)
	for l < l1 && s1[l] == s2[l] {
		l++
	}
	for r < l1 && s1[l1-1-r] == s2[l2-1-r] {
		r++
	}
	fmt.Println(l, r)
	return l+r >= l1
}
