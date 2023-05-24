package main

import (
	"bytes"
	"strconv"
)

func main() {

}

func fractionToDecimal(numerator int, denominator int) string {
	b := bytes.Buffer{}
	if (numerator * denominator) < 0 {
		b.WriteByte('-')
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	b.WriteString(strconv.Itoa(numerator / denominator))
	numerator %= denominator
	if numerator == 0 {
		return b.String()
	}
	b.WriteString(".")
	cache := make(map[int]int)
	cycleStart := -1
	for numerator > 0 {
		numerator *= 10
		if v, ok := cache[numerator]; ok {
			cycleStart = v
			break
		}
		cache[numerator] = b.Len()
		b.WriteString(strconv.Itoa(numerator / denominator))
		numerator %= denominator
	}
	ans := b.String()
	if cycleStart > 0 {
		ans = ans[:cycleStart] + "(" + ans[cycleStart:] + ")"
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
