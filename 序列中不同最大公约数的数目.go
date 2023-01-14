package main

func main() {

}

func countDifferentSubsequenceGCDs(nums []int) int {
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	cache := make([]bool, max+1)
	for _, num := range nums {
		cache[num] = true
	}
	ans := 0
	for i := 1; i <= max; i++ {
		g := 0
		for j := i; j <= max && g != i; j += i {
			if cache[j] {
				g = gcd(g, j)
			}
		}
		if g == i {
			ans++
		}
	}
	return ans
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}
