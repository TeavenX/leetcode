package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(devide(-2147483648, 2147483647))
}

func devide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	// 除数绝对值最大，只存在1（两数相等）和0（两数不相等）两种情况，特殊处理防止绝对值溢出
	if b == math.MinInt32 {
		if a == b {
			return 1
		} else {
			return 0
		}
	}

	sign := 1
	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		sign = -1
	}

	a = intAbs(a)
	b = intAbs(b)

	res := 0

	for i := 32; i >= 0; i-- {
		// 不直接用`>b`的原因 ?
		if (a>>i)-b >= 0 {
			a = a - (b << i)
			res += 1 << i
		}
	}
	return res * sign

}

func intAbs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
