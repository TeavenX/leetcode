package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cpddomains := []string{"9001 discuss.leetcode.com"}
	fmt.Println(subdomainVisits(cpddomains))
}

func subdomainVisits(cpdomains []string) []string {
	cache := make(map[string]int)
	for _, cpddomain := range cpdomains {
		t := strings.Split(cpddomain, " ")
		times, _ := strconv.Atoi(t[0])
		domain := t[1]
		cache[domain] += times
		dots := strings.Count(cpddomain, ".")
		for i := 0; i < dots; i++ {
			cpddomain = strings.SplitN(cpddomain, ".", 2)[1]
			cache[cpddomain] += times
		}
	}
	result := make([]string, 0, len(cache))
	for k, v := range cache {
		result = append(result, strings.Join([]string{strconv.Itoa(v), k}, " "))
	}
	return result
}
