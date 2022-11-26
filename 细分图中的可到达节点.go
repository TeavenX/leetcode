package main

import (
	"container/heap"
	"math"
)

func main() {

}

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	result := 0
	graph := make([][]neighbor, n)
	for _, edge := range edges { // 建图
		u, v, dist := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], neighbor{to: v, weight: dist + 1})
		graph[v] = append(graph[v], neighbor{to: u, weight: dist + 1})
	}
	dist := dijkstra(graph, 0) // 算出从0出发到各个点的最短路径

	for _, d := range dist {
		if d <= maxMoves {
			result++
		}
	}
	for _, edge := range edges {
		u, v, cnt := edge[0], edge[1], edge[2]
		a := max(maxMoves-dist[u], 0)
		b := max(maxMoves-dist[v], 0)
		result += min(a+b, cnt)
	}
	return result
}

type neighbor struct {
	to, weight int
}

func dijkstra(graph [][]neighbor, start int) []int {
	dist := make([]int, len(graph))
	for i := range graph {
		dist[i] = math.MaxInt
	}
	dist[start] = 0 // 初始化最小距离切片
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		if dist[x] < p.dist { // 距离变短了才更新，没有变短跳过
			continue
		}
		for _, e := range graph[x] {
			y := e.to
			if d := dist[x] + e.weight; d < dist[y] { // 经过x点的距离如果比原距离要更小，就更新原距离
				dist[y] = d
				heap.Push(&h, pair{y, d}) // 更新后，经过y点的距离也会发生变化，需要push回去更新
			}
		}
	}
	return dist
}

type pair struct {
	x, dist int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() interface{} {
	r := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
