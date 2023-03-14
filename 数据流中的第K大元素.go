package main

import "container/heap"

func main() {

}

type HP []int

func (h HP) Len() int {
	return len(h)
}

func (h HP) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HP) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func (h HP) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *HP) Push(v interface{}) {
	*h = append(*h, v.(int))
}

type KthLargest struct {
	k    int
	data HP
}

func Constructor(k int, nums []int) KthLargest {
	data := HP{}
	for _, num := range nums {
		heap.Push(&data, num)
		if len(data) > k {
			heap.Pop(&data)
		}
	}
	return KthLargest{
		k:    k,
		data: data,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&(this.data), val)
	if len(this.data) > this.k {
		heap.Pop(&(this.data))
	}
	return this.data[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
