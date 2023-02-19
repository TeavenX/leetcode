package main

import "container/heap"

func main() {

}

type pair struct {
	a, n int
	dta  float64
}

type Heap []pair

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i].dta > (*h)[j].dta
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *Heap) Pop() interface{} {
	top := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return top
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	var ans float64
	h := &Heap{}
	for _, cls := range classes {
		n, a := cls[1], cls[0]
		heap.Push(h, pair{a: a, n: n, dta: float64(n-a) / float64(n*(n+1))})
	}
	for i := 0; i < extraStudents; i++ {
		t := heap.Pop(h).(pair)
		p := pair{n: t.n + 1, a: t.a + 1}
		p.dta = float64(p.n-p.a) / float64(p.n*(p.n+1))
		heap.Push(h, p)
	}
	for h.Len() > 0 {
		t := heap.Pop(h).(pair)
		ans += float64(t.a) / float64(t.n)
	}
	return ans / float64(len(classes))
}
