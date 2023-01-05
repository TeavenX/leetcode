package main

func main() {

}

func countEven(num int) int {
	count := 0
	for i := 2; i <= num; i++ {
		n := i
		t := 0
		for n > 0 {
			t += n % 10
			n /= 10
		}
		if t&1 == 0 {
			count++
		}
	}
	return count
}
