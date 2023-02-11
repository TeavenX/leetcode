package main

import "sort"

func main() {

}

func fillCups(amount []int) int {
	sort.Ints(amount)
	if amount[2] >= amount[0]+amount[1] {
		return amount[2]
	}
	// 当x + y > z时，总能实现两两匹配，最多最后剩1个
	return (amount[0] + amount[1] + amount[2] + 1) >> 1
}
