package main

import "math"

func main() {

}

func reachNumber(target int) int {
	target = abs(target)
	i := math.Sqrt(float64(8*target+1)-1) / 2
	i = float64(int(i))
	// 最小值就是一直朝一个方向走，等差数列求和
	if i*(i+1) == float64(2*target) {
		return int(i)
	}

	idx := int(i)
	sum := (1 + idx) * idx / 2
	// 当需要往方向走时，一个数变成负数，sum -= 2*num，sum和target的差值一定是偶数，不然走不到
	// 当第一次sum和target的差值变成偶数时，只要其中一个num变成负数就是target了
	for sum < target || (sum-target)&1 == 1 {
		idx++
		sum += idx
	}
	return idx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
