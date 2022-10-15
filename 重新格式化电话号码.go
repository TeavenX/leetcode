package main

import "strings"

func main() {

}

func reformatNumber(number string) string {
	t := make([]byte, 0)
	for _, num := range number {
		if num < '0' || num > '9' {
			continue
		}
		t = append(t, byte(num))
	}
	n := len(t)
	times := n / 3
	rest := n % 3
	if rest == 1 {
		times--
	}
	result := make([]string, 0, times+1)
	i := 0
	for ; i < times; i++ {
		result = append(result, string(t[i*3:(i+1)*3]))
	}
	if rest == 1 || rest == 2 {
		for idx := i * 3; idx < n; idx += 2 {
			result = append(result, string(t[idx:idx+2]))
		}
	}
	return strings.Join(result, "-")
}
