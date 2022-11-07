package main

import "fmt"

func main() {

}

func ambiguousCoordinatesError(s string) []string {
	result := make([]string, 0)
	left := make(map[int][]string)
	right := make(map[int][]string)
	s = s[1 : len(s)-1]
	n := len(s)
	var traceleft func(idx int, zeroCount int, hasdiv bool, str string)
	var traceright func(idx int, hasdiv bool, str string)
	traceleft = func(idx int, zeroCount int, hasdiv bool, str string) {
		if idx > 0 && (zeroCount != idx || zeroCount == 1) && (!hasdiv || str[len(str)-1] != '0') {
			left[idx] = append(left[idx], str)
		}
		if idx == n-1 {
			return
		}
		if zeroCount == 1 && !hasdiv {
			if s[idx] == '0' {
				zeroCount++
			}
			traceleft(idx+1, zeroCount, true, str+"."+s[idx:idx+1])
		} else {
			if s[idx] == '0' {
				zeroCount++
			}
			traceleft(idx+1, zeroCount, hasdiv, str+s[idx:idx+1])
			if !hasdiv && len(str) > 0 {
				traceleft(idx+1, zeroCount, true, str+"."+s[idx:idx+1])
			}
		}
	}
	traceright = func(idx int, hasdiv bool, str string) {
		if idx < n-1 && (len(str) == 1 || str[0] != '0' || (str[1] == '.' && str[len(str)-1] != '0')) && (!hasdiv || str[len(str)-1] != '0') {
			right[idx+1] = append(right[idx+1], str)
		}
		if idx == 0 {
			return
		}
		traceright(idx-1, hasdiv, s[idx:idx+1]+str)
		if !hasdiv && len(str) > 0 {
			traceright(idx-1, true, s[idx:idx+1]+"."+str)
		}
	}
	traceleft(0, 0, false, "")
	traceright(n-1, false, "")
	for i := 1; i <= n; i++ {
		for _, l := range left[i] {
			for _, r := range right[i] {
				result = append(result, fmt.Sprintf("(%s, %s)", l, r))
			}
		}
	}
	return result
}
