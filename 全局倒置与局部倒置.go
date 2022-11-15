package main

func main() {

}

func isIdealPermutation(nums []int) bool {
	// 满足局部倒置一定满足全局倒置
	// nums 是范围 [0, n - 1] 内所有数字组成的一个排列
	// 当一个数和排序后的nums下标超过1，就不符合规范
	for i := 0; i < len(nums); i++ {
		if abs(nums[i]-i) > 1 {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
