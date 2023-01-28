package main

func main() {

}

func countAsterisks(s string) int {
	flag := true
	ans := 0
	for _, b := range s {
		if b == '|' {
			flag = !flag
		}
		if flag && b == '*' {
			ans++
		}
	}
	return ans
}
