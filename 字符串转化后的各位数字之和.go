package main

func main() {

}

func getLucky(s string, k int) int {
	var result rune = 0
	for _, b := range s {
		num := b - 'a' + 1
		result += num/10 + num%10
	}
	for i := 1; i < k; i++ {
		var temp rune = 0
		for result > 0 {
			temp += result % 10
			result /= 10
		}
		result = temp
	}
	return int(result)
}
