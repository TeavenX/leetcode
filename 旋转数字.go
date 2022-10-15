package main

func main() {

}

func rotatedDigits(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		ok, x := false, i
		for x > 0 {
			t := x % 10
			x /= 10
			if t == 2 || t == 5 || t == 6 || t == 9 {
				ok = true
			} else if t != 0 && t != 1 && t != 8 {
				ok = false
				break
			}
		}
		if ok {
			result++
		}
	}
	return result
}

func rotatedDigitsV2(n int) int {
	//plus := []int{0,1,2,5,6,8,9}
	result := 0
	cache := map[int]bool{0: true, 1: true, 2: true, 5: true, 6: true, 8: true, 9: true}
	for i := 1; i <= n; i++ {
		b, s := i/10, i%10
		if cache[b] && cache[s] {
			result++
			cache[i] = true
		}
	}
	blacklist := []int{1, 8, 10, 11, 18}
	for _, num := range blacklist {
		if num <= n {
			result--
		}
	}
	return result
}
