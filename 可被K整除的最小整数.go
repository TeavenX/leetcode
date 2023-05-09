package main

func main() {

}

func smallestRepunitDivByK(k int) int {
	if k&1 == 0 || k%5 == 0 {
		return -1
	}
	n := 0
	for i := 1; i <= k; i++ {
		n = (n*10 + 1) % k
		if n%k == 0 {
			return i
		}
	}
	return -1
}
