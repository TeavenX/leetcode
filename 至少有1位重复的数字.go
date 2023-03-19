package main

import "strconv"

func main() {

}

func numDupDigitsAtMostN(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	cache := make([][1 << 10]int, m) // 一共有10个数字可以选择
	for i := range cache {
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var search func(i int, mask int, isLimit bool, isNum bool) int
	// i 表示递归的深度
	// mask 用位图记录每一个数字的使用情况, mask 上限是 1 << 10，一共有10个数字可以选择
	// isLimit 表示当前深度选择数字的时候，是否有上限，例如数字123，深度是2的时候，也就是选择数字2，
	//         如果前面一位数字已经选择了1，这里 limit 就是2，而不是9；反之，如果前面选择的是0，这里limit就是9
	// isNum 表示当前是否已经构成了一个数字。这里可以用 mask > 0 来判断，当 mask > 0 时表示前面已经有选择过数字了
	// 返回深度为 i 时，当前 mask 一共可以满足多少种结果
	search = func(i int, mask int, isLimit bool, isNum bool) int {
		if i == m {
			// 递归到了最深一层
			if isNum {
				return 1
			}
			return 0
		}
		if v := cache[i][mask]; !isLimit && isNum && v > 0 {
			// 这里要判断 isLimit 和 isNum 的原因是：下面的计算中统计的是期望的结果
			// 也就是 无论是否有限制 以及 能构成数字
			return v
		}
		res := 0
		if !isNum {
			// 前面都没选数字，可以直接到下一层
			res += search(i+1, mask, false, false)
		}
		start := 0
		if !isNum {
			// 前面的没选择数字，为了保证数字合法，需要从1开始选择数字
			start = 1
		}
		limit := 9
		if isLimit {
			// 如果前面都是往最大取的话，这里最高不能超过原因数字位中的值
			limit = int(s[i] - '0')
		}
		for j := start; j <= limit; j++ {
			if mask>>j&1 == 0 {
				// 数字没有选过
				res += search(i+1, mask|1<<j, isLimit && j == limit, true)
			}
		}
		cache[i][mask] = res
		return res
	}
	return n - search(0, 0, true, false)
}
