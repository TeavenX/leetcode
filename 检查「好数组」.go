package main

func main() {

}

func isGoodArray(nums []int) bool {
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		r = gcd(r, nums[i])
		if r == 1 {
			return true
		}
	}
	return r == 1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// func gcd(a, b int) int {
//     if b == 0 {
//         return a
//     }
//     return gcd(b, a % b)
// }
