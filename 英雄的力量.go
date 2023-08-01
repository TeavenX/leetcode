package main

import "sort"

func main() {

}

/*
排序后的数组是：a, b, c, d, e
考虑 d 作为最大值有以下几种选择
1.只包含d - d^2 * d
2.至少包含a和d（b和c可选可不选） - d^2 * a * 2^2(都不选；都选；只选b；只选c；四种可能但是最大最小值不变)
3.至少包含b和d（c可选可不选） - d^2 * b * 2^1(选c；不选c)
4.至少包含c和d - d^2 * c * 2^0
sum = d^3 + d^2 * (a * 2^2 + b * 2^1 + c * 2^0)
令 s = (a * 2^2 + b * 2^1 + c * 2^0)
sum = d^2 * (d + s)

同理考虑 e 最为最大值
sum = e^3 + e^2 * (a * 2^3 + b * 2^2 + c * 2^1 + d * 2^0)
    = e^3 + e^2 * (2 * (a * 2^2 + b * 2^1 + c * 2^0) + d * 2^0 )
    = e^3 + e^2 * (2 * s + d)
    = e^2 * (e + (2 * s + d))
此时 s = 2 * s + d
*/

const mod int = 1e9 + 7

func sumOfPower(nums []int) int {
	sort.Ints(nums)
	s := 0
	ans := 0
	for _, x := range nums {
		ans = (ans + (x*x)%mod*(x+s)) % mod
		s = (s*2 + x) % mod
	}
	return ans
}
