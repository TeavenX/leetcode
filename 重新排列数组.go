package main

func main() {

}

func shuffle(nums []int, n int) []int {
	result := make([]int, 2*n)
	for i, j := 0, 0; i < n; i++ {
		result[j] = nums[i]
		result[j+1] = nums[n+i]
		j += 2
	}
	return result
}

// int32一共有32位，而题目给出num的最大值是1000，2^10=1024，也就是实际只需要用到32位中的10位
// 因此可以用末10位存当前的值，20-10位存shuffle之后的值，最后把末10未抹掉就是期望的答案
// 此方法实现了O(1)空间复杂度
func shuffleV2(nums []int, n int) []int {
	for i := 0; i < n; i++ {
		nums[2*i] |= (nums[i] & 1023) << 10
		nums[2*i+1] |= (nums[i+n] & 1023) << 10
	}
	for i := 0; i < 2*n; i++ {
		nums[i] = nums[i] >> 10
	}
	return nums
}
