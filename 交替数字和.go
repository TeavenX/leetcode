package main

func main() {

}

func alternateDigitSum(n int) int {
	ans := 0
	flag := 1
	for n > 0 {
		ans += flag * (n % 10)
		n /= 10
		flag = -flag
	}
	if flag > 0 {
		return -ans
	}
	return ans
}
