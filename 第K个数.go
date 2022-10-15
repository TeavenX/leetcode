package main

func main() {

}

func getKthMagicNumberError(k int) int {
	// 3*3*3 < 5*7
	if k <= 0 {
		return 0
	}
	m := []int{3, 5, 7}
	if k == 1 {
		return 1
	} else if k < 5 {
		return m[k-2]
	}
	k -= 5
	times := k / 6
	idx := k % 6
	arr := []int{9, 15, 21, 25, 35, 49}
	for i := 0; i < times; i++ {
		arr[idx] *= m[i%3]
	}
	return arr[idx]
}

func getKthMagicNumber(k int) int {
	if k <= 0 {
		return 0
	}
	ugly := make([]int, k)
	ugly[0] = 1
	p3, p5, p7 := 0, 0, 0
	for i := 1; i < k; i++ {
		ugly[i] = min(ugly[p3]*3, min(ugly[p5]*5, ugly[p7]*7))
		if ugly[i] == ugly[p3]*3 {
			p3++
		}
		if ugly[i] == ugly[p5]*5 {
			p5++
		}
		if ugly[i] == ugly[p7]*7 {
			p7++
		}
	}
	return ugly[k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
