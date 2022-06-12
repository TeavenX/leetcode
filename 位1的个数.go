package main

func main() {

}

func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		num = num & (num - 1)
		result++
	}
	return result
}

func hammingWeight20220613(num uint32) int {
	result := 0
	for num > 0 {
		num = (num - 1) & num
		result++
	}
	return result
}
