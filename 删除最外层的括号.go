package main

func main() {

}

func removeOuterParentheses(s string) string {
	result := ""
	pre, flag := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			flag++
		} else {
			flag--
		}
		if flag == 0 && i > 0 {
			result += s[pre+1 : i]
			pre = i + 1
		}
	}
	return result
}

func removeOuterParenthesesV2(s string) string {
	result := make([]rune, 0)
	flag := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			flag--
		}
		if flag > 0 {
			result = append(result, rune(s[i]))
		}
		if s[i] == '(' {
			flag++
		}
	}
	return string(result)
}
