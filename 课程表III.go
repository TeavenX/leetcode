package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1]
	})
	cur := 0
	h := &HP{}
	for _, c := range courses {
		d := c[0]
		if cur+d <= c[1] {
			cur += d
			heap.Push(h, d)
		} else if h.Len() > 0 && d < h.IntSlice[0] {
			// 如果当前的时间，比之前最长的耗时要短，就置换，但是课程数量不变
			cur -= h.replace(d) - d
		}
	}
	return h.Len()
}

type HP struct {
	sort.IntSlice
}

func (h *HP) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *HP) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *HP) Pop() interface{} {
	t := h.IntSlice
	v := t[len(t)-1]
	h.IntSlice = t[:len(t)-1]
	return v
}

func (h *HP) replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(h, 0)
	return top
}
