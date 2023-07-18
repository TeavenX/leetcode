package main

import (
	"container/heap"
	"sort"
)

func main() {

}

type Pair struct {
	v int
	r int
}

type HP []Pair

func (h HP) Len() int {
	return len(h)
}

func (h HP) Less(i, j int) bool {
	return h[i].v < h[j].v
}

func (h HP) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HP) Push(v interface{}) {
	*h = append(*h, v.(Pair))
}

func (h *HP) Pop() interface{} {
	t := *h
	v := t[len(t)-1]
	t = t[:len(t)-1]
	*h = t
	return v
}

func minInterval(intervals [][]int, queries []int) []int {
	n, m := len(intervals), len(queries)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	sortedQueries := make([][2]int, m)
	ans := make([]int, m)
	for i, q := range queries {
		sortedQueries[i] = [2]int{q, i}
		ans[i] = -1
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][0] < sortedQueries[j][0]
	})
	idx := 0
	hp := HP{}
	for _, q := range sortedQueries {
		for idx < n && intervals[idx][0] <= q[0] {
			l, r := intervals[idx][0], intervals[idx][1]
			heap.Push(&hp, Pair{r - l + 1, r})
			idx++
		}
		for len(hp) > 0 && hp[0].r < q[0] {
			heap.Pop(&hp)
		}
		if len(hp) > 0 {
			ans[q[1]] = hp[0].v
		}
	}
	return ans
}
