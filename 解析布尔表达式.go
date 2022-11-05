package main

import "fmt"

func main() {
	fmt.Println(parseBoolExpr("|(f,f,f,t)"))
}

func parseBoolExpr(expression string) bool {
	if expression == "t" {
		return true
	}
	if expression == "f" {
		return false
	}
	parseSub := func(i int) (int, bool) {
		start := i
		c := 1
		for end := i + 2; end < len(expression); end++ {
			if expression[end] == '(' {
				c++
			} else if expression[end] == ')' {
				c--
			}
			if c == 0 {
				return end, parseBoolExpr(expression[start : end+1])
			}
		}
		return start, false
	}
	switch expression[0] {
	case '!':
		return !parseBoolExpr(expression[2 : len(expression)-1])
	case '&':
		r := true
		for i := 2; i < len(expression)-1; i++ {
			if expression[i] != 't' && expression[i] != 'f' && expression[i] != ',' {
				idx, v := parseSub(i)
				r = r && v
				i = idx
			} else if expression[i] == 'f' {
				return false
			}
		}
		return r
	case '|':
		r := false
		for i := 2; i < len(expression)-1; i++ {
			if expression[i] != 't' && expression[i] != 'f' && expression[i] != ',' {
				idx, v := parseSub(i)
				r = r || v
				i = idx
			} else if expression[i] == 't' {
				return true
			}
		}
		return r
	}
	return false
}
