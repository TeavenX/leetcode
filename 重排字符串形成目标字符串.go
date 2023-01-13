package main

import "math"

func main() {

}

func rearrangeCharacters(s string, target string) int {
	if len(s) < len(target) {
		return 0
	}
	cache := [26]int{}
	for _, b := range target {
		cache[b-'a']++
	}
	temp := [26]int{}
	ans := math.MaxInt
	for _, b := range s {
		temp[b-'a']++
	}
	for k, v := range cache {
		if v > 0 {
			if t := temp[k] / v; t < ans {
				ans = t
			}
		}
	}
	return ans
}
