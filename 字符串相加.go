package main

func main() {

}

func addStrings(num1 string, num2 string) string {
	ans := make([]byte, max(len(num1), len(num2))+1)
	idx := len(ans) - 1
	var carry byte
	for len(num1) > 0 && len(num2) > 0 {
		t := num1[len(num1)-1] - '0' + num2[len(num2)-1] - '0' + carry
		num1 = num1[:len(num1)-1]
		num2 = num2[:len(num2)-1]
		carry = t / 10
		t %= 10
		ans[idx] = t + '0'
		idx--
	}
	for len(num1) > 0 {
		t := num1[len(num1)-1] - '0' + carry
		num1 = num1[:len(num1)-1]
		carry = t / 10
		t %= 10
		ans[idx] = t + '0'
		idx--
	}
	for len(num2) > 0 {
		t := num2[len(num2)-1] - '0' + carry
		num2 = num2[:len(num2)-1]
		carry = t / 10
		t %= 10
		ans[idx] = t + '0'
		idx--
	}
	if carry > 0 {
		ans[idx] = carry + '0'
	}
	if ans[0] == 0 {
		ans = ans[1:]
	}
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
