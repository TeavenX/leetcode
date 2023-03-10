package main

func main() {

}

func minSubarray(nums []int, p int) int {
	x := 0
	for _, num := range nums {
		x += num
	}
	x %= p
	if x == 0 {
		return 0
	}
	n := len(nums)
	ans, sum := n, 0
	last := map[int]int{sum: -1} // 这里初始化1的原因是right从0开始，如果删除第一个的话，应该是0-(-1) = 1
	// (sum[right] % p − x % p + p) % p = sum[left] % p
	for right, num := range nums {
		sum += num
		last[sum%p] = right
		if left, ok := last[(sum-x+p)%p]; ok {
			ans = min(ans, right-left)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
