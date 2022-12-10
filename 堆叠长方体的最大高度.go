package main

import "sort"

func main() {

}

func maxHeight(cuboids [][]int) int {
	n := len(cuboids)
	for _, cuboid := range cuboids {
		sort.Ints(cuboid)
	}
	sort.Slice(cuboids, func(i, j int) bool {
		a, b := cuboids[i], cuboids[j]
		return a[0] > b[0] || a[0] == b[0] && (a[1] > b[1] || a[1] == b[1] && a[2] > b[2])
	})
	dp := make([]int, n) // dp[i]表示第i个长方体所能堆到的最大高度，可以堆到[0, i-1]的任何一个长方体上
	result := 0
	for i, c1 := range cuboids {
		for j, c2 := range cuboids[:i] {
			if c1[1] <= c2[1] && c1[2] <= c2[2] { // c1[0] <= c2[0]恒成立（前面排序过了）
				dp[i] = max(dp[i], dp[j]) // dp[i]会被历史的大值覆盖，只要满足能堆叠，就取最大的
			}
		}
		dp[i] += c1[2]              // 加上当前的高度
		result = max(result, dp[i]) // 遍历的最后一个不一定是最高的，所以这里每次都取最大值
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
