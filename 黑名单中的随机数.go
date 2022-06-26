package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	blacklist := []int{3, 6, 10}
	s := Constructor(11, blacklist)
	fmt.Println(s.Pick(), s.Pick(), s.Pick(), s.Pick(), s.Pick())
}

type Solution struct {
	n          int
	projection map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	sort.Ints(blacklist)
	projection := map[int]int{}
	m := len(blacklist)
	for _, bn := range blacklist {
		projection[bn] = -1
	}
	bidx := n - 1
	for _, bn := range blacklist {
		if bn >= n-m {
			continue
		}
		for projection[bidx] == -1 {
			bidx--
		}
		projection[bn] = bidx
		bidx--
	}
	return Solution{n: n - m, projection: projection}
}

func (this *Solution) Pick() int {
	num := rand.Intn(this.n)
	if this.projection[num] > 0 {
		return this.projection[num]
	}
	return num
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
