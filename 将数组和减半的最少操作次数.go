package main

import "container/heap"

func main() {

}

type HP []float64

func (h HP) Len() int {
	return len(h)
}

func (h HP) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h HP) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HP) Push(v interface{}) {
	*h = append(*h, v.(float64))
}

func (h *HP) Pop() interface{} {
	t := *h
	v := t[len(t)-1]
	t = t[:len(t)-1]
	*h = t
	return v
}

func halveArray(nums []int) int {
	var sum, cur float64
	hp := HP{}
	for _, num := range nums {
		sum += float64(num)
		heap.Push(&hp, float64(num))
	}
	cur = sum
	sum /= 2
	ans := 0
	for cur > sum {
		v := heap.Pop(&hp).(float64)
		v /= 2
		cur -= v
		heap.Push(&hp, v)
		ans++
	}
	return ans
}
