package main

import "sort"

func main() {

}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	index := make([]int, n)
	for i := range index {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		ii, jj := index[i], index[j]
		return reward1[ii]-reward2[ii] > reward1[jj]-reward2[jj]
	})
	ans := 0
	blk := make(map[int]bool, k)
	for i := 0; i < k; i++ {
		idx := index[i]
		blk[idx] = true
		ans += reward1[idx]
	}
	for i, r := range reward2 {
		if !blk[i] {
			ans += r
		}
	}
	return ans
}

func miceAndCheeseV2(reward1 []int, reward2 []int, k int) int {
	ans := 0
	for i, r := range reward2 {
		ans += r
		reward1[i] -= r
	}
	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
	for _, r := range reward1[:k] {
		ans += r
	}
	return ans
}
