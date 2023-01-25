package main

import "sort"

func main() {

}

type UnionFind struct {
	parents, ranks []int
}

func NewUnionFind(n int) *UnionFind {
	p, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		r[i] = 1
	}
	return &UnionFind{
		parents: p,
		ranks:   r,
	}
}

func (u *UnionFind) find(x int) int {
	if u.parents[x] != x {
		return u.find(u.parents[x])
	}
	return u.parents[x]
}

func (u *UnionFind) union(x, y int) {
	px, py := u.find(x), u.find(y)
	if px == py {
		return
	}
	if u.ranks[px] > u.ranks[py] {
		u.parents[py] = px
		u.ranks[px] += u.ranks[py]
	} else {
		u.parents[px] = py
		u.ranks[py] += u.ranks[px]
	}
}

func (u *UnionFind) reset(x int) {
	u.parents[x] = x
	u.ranks[x] = 1
}

type pair struct {
	i, j int
}

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	cache := make(map[int][]pair)
	for i, row := range matrix {
		for j, val := range row {
			cache[val] = append(cache[val], pair{i, j})
		}
	}
	rowMax := make([]int, m)
	colMax := make([]int, n)
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	vs := make([]int, 0)
	for v := range cache {
		vs = append(vs, v)
	}
	sort.Ints(vs)
	unionFind := NewUnionFind(m + n)
	rank := make([]int, m+n)
	for _, v := range vs {
		ps := cache[v]
		for _, p := range ps {
			unionFind.union(p.i, p.j+m)
		}
		for _, p := range ps {
			i, j := p.i, p.j
			rank[unionFind.find(i)] = max(rank[unionFind.find(i)], max(rowMax[i], colMax[j]))
		}
		for _, p := range ps {
			i, j := p.i, p.j
			ans[i][j] = 1 + rank[unionFind.find(i)]
			rowMax[i], colMax[j] = ans[i][j], ans[i][j]
		}
		for _, p := range ps {
			unionFind.reset(p.i)
			unionFind.reset(p.j + m)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
