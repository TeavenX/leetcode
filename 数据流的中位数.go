package main

import (
	"container/heap"
)

func main() {

}

type MedianFinder struct {
	mx *HP
	mn *HP
}

type HP struct {
	data []int
	mx   bool
}

func (h *HP) Len() int {
	return len(h.data)
}

func (h *HP) Less(i, j int) bool {
	if h.mx {
		return h.data[i] < h.data[j]
	}
	return h.data[i] > h.data[j]
}

func (h *HP) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *HP) Push(v interface{}) {
	h.data = append(h.data, v.(int))
}

func (h *HP) Pop() interface{} {
	v := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return v
}

func Constructor() MedianFinder {
	return MedianFinder{
		mx: &HP{mx: true},
		mn: &HP{mx: false},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.mn.Len() > this.mx.Len() {
		heap.Push(this.mn, num)
		num = heap.Pop(this.mn).(int)
		heap.Push(this.mx, num)
	} else {
		heap.Push(this.mx, num)
		num = heap.Pop(this.mx).(int)
		heap.Push(this.mn, num)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	count := this.mn.Len() + this.mx.Len()
	if count&1 == 1 {
		return float64(this.mn.data[0])
	}
	return float64(this.mn.data[0]+this.mx.data[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
