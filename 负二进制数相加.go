package main

func main() {

}

func addNegabinary(arr1 []int, arr2 []int) []int {
	n1 := len(arr1)
	n2 := len(arr2)
	n := max(n1, n2) + 4
	ans := make([]int, n)
	for i := n1 - 1; i >= 0; i-- {
		ans[n1-1-i] += arr1[i]
	}
	for i := n2 - 1; i >= 0; i-- {
		ans[n2-1-i] += arr2[i]
	}
	for i := 0; i+2 < n; i++ {
		carry := ans[i] >> 1
		ans[i] &= 1
		ans[i+1] += carry
		ans[i+2] += carry
	}
	k := n - 3 // 最后两位是纯进位，不需要管（上限是n+2的长度），这里原本是从n-1开始，减去最后两位是n-3
	for k > 0 && ans[k] == 0 {
		// 去除前导0
		k--
	}
	for i := 0; i < (k+1)>>1; i++ {
		ans[i], ans[k-i] = ans[k-i], ans[i]
	}
	return ans[:k+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
