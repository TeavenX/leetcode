package main

func main() {

}

func maxScore(s string) int {
	count1 := 0
	for _, num := range s {
		if num == '1' {
			count1++
		}
	}
	if count1 == len(s) || count1 == 0 {
		return len(s) - 1
	}
	max := 0
	cur := 0
	for i, num := range s {
		if num == '0' {
			cur++
		} else {
			count1--
		}
		if cur+count1 > max && i < len(s) {
			max = cur + count1
		}
	}
	return max
}
