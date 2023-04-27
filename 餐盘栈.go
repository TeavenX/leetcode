package main

import "container/heap"

func main() {

}

type HP []int

func (h HP) Len() int {
	return len(h)
}

func (h HP) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h HP) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HP) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *HP) Pop() interface{} {
	t := *h
	v := t[len(t)-1]
	t = t[:len(t)-1]
	*h = t
	return v
}

type DinnerPlates struct {
	capacity int
	hp       HP
	stacks   [][]int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		capacity: capacity,
		hp:       HP{},
		stacks:   make([][]int, 0),
	}
}

func (d *DinnerPlates) Push(val int) {
	if d.hp.Len() > 0 && d.hp[0] >= len(d.stacks) {
		d.hp = HP{}
	}
	if d.hp.Len() == 0 {
		d.stacks = append(d.stacks, []int{val})
		if d.capacity > 1 {
			heap.Push(&d.hp, len(d.stacks)-1)
		}
	} else {
		i := d.hp[0]
		d.stacks[i] = append(d.stacks[i], val)
		if len(d.stacks[i]) == d.capacity {
			heap.Pop(&d.hp)
		}
	}
}

func (d *DinnerPlates) Pop() int {
	return d.PopAtStack(len(d.stacks) - 1)
}

func (d *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(d.stacks) || len(d.stacks[index]) == 0 {
		return -1
	}
	if len(d.stacks[index]) == d.capacity {
		heap.Push(&d.hp, index)
	}
	idx := len(d.stacks[index]) - 1
	val := d.stacks[index][idx]
	d.stacks[index] = d.stacks[index][:idx]
	for len(d.stacks) > 0 && len(d.stacks[len(d.stacks)-1]) == 0 {
		d.stacks = d.stacks[:len(d.stacks)-1]
	}
	return val
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
