package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//stations := [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}
	stations := [][]int{{10, 20}, {20, 60}, {30, 10}, {80, 40}}
	fmt.Println(minRefuelStopsV2(100, 10, stations))
	stations = [][]int{{25, 50}, {50, 25}}
	fmt.Println(minRefuelStopsV2(100, 50, stations))
}

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	if startFuel >= target {
		return 0
	}
	if len(stations) == 0 || stations[0][0] > startFuel {
		return -1
	}
	count := 0
	maxHeap := MaxHeap{}
	start := 0
	for _, station := range stations {
		for startFuel < station[0]-start {
			if len(maxHeap) == 0 {
				return -1
			}
			startFuel += heap.Pop(&maxHeap).(int)
			count++
		}
		heap.Push(&maxHeap, station[1])
		startFuel -= station[0] - start
		start = station[0]
		if start+startFuel >= target {
			return count
		}
	}
	for startFuel < target-start {
		if len(maxHeap) == 0 {
			return -1
		}
		startFuel += heap.Pop(&maxHeap).(int)
		count++
	}
	return count
}

func minRefuelStopsV2(target int, startFuel int, stations [][]int) int {
	if startFuel >= target {
		return 0
	}
	if len(stations) == 0 || stations[0][0] > startFuel {
		return -1
	}
	count := 0
	maxHeap := MaxHeap{}
	for _, station := range stations {
		for startFuel < station[0] {
			if len(maxHeap) == 0 {
				return -1
			}
			startFuel += heap.Pop(&maxHeap).(int)
			count++
		}
		heap.Push(&maxHeap, station[1])
	}
	for startFuel < target {
		if len(maxHeap) == 0 {
			return -1
		}
		startFuel += heap.Pop(&maxHeap).(int)
		count++
	}
	return count
}
