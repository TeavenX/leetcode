package main

import "sort"

func main() {

}

func maximumTastiness(price []int, k int) int {
	n := len(price)
	sort.Ints(price)
	if k == 2 {
		return price[n-1] - price[0]
	}
	if price[0] == price[n-1] {
		return 0
	}
	find := func(d int) int {
		cnt, pre := 1, price[0]
		for _, p := range price {
			if p >= pre+d {
				cnt++
				pre = p
			}
		}
		return cnt
	}
	left, right := 0, (price[n-1]-price[0])/(k-1)+1 // (0, (max-min)/(k-1) + 1) 开区间向下取整
	for left+1 < right {                            // 由于是开区间，left和right中间必须有一个值，如果left = right，那就等于没有值了
		mid := left + (right-left)>>1
		if find(mid) >= k {
			left = mid // (mid, right)
		} else {
			right = mid // (left, mid)
		}
	}
	return left
	// left, right := 1, (price[n-1] - price[0] + k-2) / (k-1) // [1, (max-min+k-1-1) / (k-1)]
	// for left <= right { // 闭区间，当 left = right 时，就是刚好剩下一个元素
	// 	mid := left + (right-left) >> 1
	// 	if find(mid) >= k {
	// 		left = mid+1 // [mid+1, right]
	// 	} else {
	// 		right = mid-1 // [left, mid-1]
	// 	}
	// }
	// return right
}
