package main

func main() {

}

func subtractProductAndSum(n int) int {
	sum, mult := 0, 1
	for ; n > 0; n /= 10 {
		num := n % 10
		sum += num
		mult *= num
	}
	return mult - sum
}
