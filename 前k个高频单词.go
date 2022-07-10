package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func topKFrequent(words []string, k int) []string {
	cache := make(map[string]int)
	for _, word := range words {
		cache[word]++
	}
	result := make([]string, 0, len(cache))
	for k := range cache {
		result = append(result, k)
	}
	sort.Slice(result, func(i, j int) bool {
		return cache[result[i]] > cache[result[j]] || (cache[result[i]] == cache[result[j]] && result[i] < result[j])
	})
	return result[:k]
}

type node struct {
	val   string
	count int
}

type hp []node

func (h hp) Less(i, j int) bool {
	return h[i].count < h[j].count || (h[i].count == h[j].count && h[i].val > h[j].val)
}

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(node))
}

func (h *hp) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func topKFrequentV2(words []string, k int) []string {
	cache := make(map[string]int)
	for _, word := range words {
		cache[word]++
	}
	h := &hp{}
	heap.Init(h)
	for word, count := range cache {
		heap.Push(h, node{val: word, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(node).val
	}
	return result
}
