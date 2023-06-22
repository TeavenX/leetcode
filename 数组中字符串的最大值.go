package main

func main() {

}

func maximumValue(strs []string) int {
	ans := 0
	for _, s := range strs {
		res := 0
		num := 0
		isNum := true
		for _, b := range s {
			if isNum && b >= '0' && b <= '9' {
				num = num*10 + int(b-'0')
			} else {
				isNum = false
			}
			res++
		}
		if isNum {
			res = num
		}
		ans = max(ans, res)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
