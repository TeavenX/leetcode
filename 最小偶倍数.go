package main

func main() {

}

func smallestEvenMultiple(n int) int {
	if n&1 == 1 {
		return 2 * n
	}
	return n
}

func smallestEvenMultipleV2(n int) int {
	return n << (n & 1)
}
