package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//stones := []int{2, 7, 4, 1, 8, 1}
	stones := []int{7, 6, 7, 6, 9}
	fmt.Println(lastStoneWeight(stones))
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) push(v int) {
	heap.Push(h, v)
}

func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

func lastStoneWeight(stones []int) int {
	n := len(stones)
	if n == 1 {
		return stones[0]
	}
	h := &hp{stones}
	heap.Init(h)
	for h.Len() > 1 {
		x, y := h.pop(), h.pop()
		if x > y {
			h.push(x - y)
		}
	}
	if h.Len() > 0 {
		return h.IntSlice[0]
	}
	return 0
}
