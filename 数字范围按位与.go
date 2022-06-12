package main

func main() {

}

func rangeBitwiseAnd(left int, right int) int {
	temp := left
	n1 := 1
	for temp > 0 {
		temp >>= 1
		n1 <<= 1
	}
	n := right - left
	if n >= n1 || left == 0 || right == 0 {
		return 0
	}
	result := left
	for i := 1; i <= n; i++ {
		result &= left + i
	}
	return result
}

func rangeBitwiseAndV2(left int, right int) int {
	if left == 0 || right == 0 {
		return 0
	}
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return right << shift
}

func rangeBitwiseAndV3(left int, right int) int {
	for left < right {
		right &= right - 1
	}
	return right
}
