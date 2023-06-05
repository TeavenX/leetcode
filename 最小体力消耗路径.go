package main

import "sort"

func main() {

}

type UnionFind struct {
	parent, rank []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
	}
}

func (f *UnionFind) find(x int) int {
	if f.parent[x] != x {
		f.parent[x] = f.find(f.parent[x])
	}
	return f.parent[x]
}

func (f *UnionFind) merge(x, y int) {
	px, py := f.find(x), f.find(y)
	if px == py {
		return
	}
	if f.rank[x] < f.rank[y] {
		px, py = py, px
	}
	f.rank[px] += f.rank[py]
	f.parent[py] = px
}

type Edge struct {
	v, w int // 边的两个端点
	diff int
}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	edges := make([]Edge, 0)
	for i, row := range heights {
		for j, h := range row {
			node := i*m + j
			if i > 0 {
				edges = append(edges, Edge{node - m, node, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, Edge{node - 1, node, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].diff < edges[j].diff
	})
	uf := NewUnionFind(n * m)
	for _, e := range edges {
		uf.merge(e.v, e.w)
		if uf.find(0) == uf.find(n*m-1) {
			return e.diff
		}
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
