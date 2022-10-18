package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func atMostNGivenDigitSet(digits []string, n int) int {
	sn := strconv.Itoa(n)
	cache := make(map[string]int)
	var search func(idx int, isLimit bool, isNum bool) int
	search = func(idx int, isLimit bool, isNum bool) int {
		if idx == len(sn) {
			if isNum {
				return 1
			}
			return 0
		}
		ci := fmt.Sprintf("%d%t%t", idx, isLimit, isNum)
		if cache[ci] > 0 {
			return cache[ci]
		}
		result := 0
		if !isNum {
			result = search(idx+1, false, false)
		}
		up := "9"
		if isLimit {
			up = string(sn[idx])
		}
		for _, d := range digits {
			if d > up {
				break
			}
			limit := isLimit && d == up
			result += search(idx+1, limit, true)
		}
		cache[ci] = result
		return result
	}
	return search(0, true, false)
}
