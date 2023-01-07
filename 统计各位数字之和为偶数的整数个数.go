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

func countEvenV2(num int) int {
	count := 0
	y, x := num/10, num%10
	count += y * 5
	ySum := 0
	for ; y > 0; y /= 10 {
		ySum += y % 10
	}
	if ySum&1 == 1 {
		count += (x+1)>>1 - 1
	} else {
		count += x >> 1
	}
	return count
}
