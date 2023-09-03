package main

import (
	"container/heap"
	"sort"
)

func main() {

}

type Pair struct {
	dist  int
	speed int
}

type HP []Pair

func (h HP) Len() int {
	return len(h)
}

func (h HP) Less(i, j int) bool {
	t1 := (h[i].dist + h[i].speed - 1) / h[i].speed
	t2 := (h[j].dist + h[j].speed - 1) / h[j].speed
	if t1 == t2 {
		return h[i].dist < h[j].dist
	}
	return t1 < t2
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

func eliminateMaximum(dist []int, speed []int) int {
	hp := &HP{}
	for i := 0; i < len(dist); i++ {
		heap.Push(hp, Pair{dist: dist[i], speed: speed[i]})
	}
	ans := 0
	cost := 0
	for hp.Len() > 0 {
		p := heap.Pop(hp).(Pair)
		if p.speed*cost >= p.dist {
			break
		}
		ans++
		cost++
	}
	return ans
}

func eliminateMaximumV2(dist []int, speed []int) int {
	n := len(dist)
	time := make([]int, n)
	for i := 0; i < n; i++ {
		time[i] = (dist[i] + speed[i] - 1) / speed[i]
	}
	sort.Ints(time)
	ans := 0
	idx := 0
	for i := 0; i < n; i++ {
		if time[idx] <= i {
			break
		}
		ans++
		idx++
	}
	return ans
}
