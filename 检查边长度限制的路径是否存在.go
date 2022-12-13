package main

import (
	"fmt"
	"sort"
)

func main() {
	edgeList := [][]int{{9, 1, 53}, {3, 2, 66}, {12, 5, 99}, {9, 7, 26}, {1, 4, 78}, {11, 1, 62}, {3, 10, 50}, {12, 1, 71}, {12, 6, 63}, {1, 10, 63}, {9, 10, 88}, {9, 11, 59}, {1, 4, 37}, {4, 2, 63}, {0, 2, 26}, {6, 12, 98}, {9, 11, 99}, {4, 5, 40}, {2, 8, 25}, {4, 2, 35}, {8, 10, 9}, {11, 9, 25}, {10, 11, 11}, {7, 6, 89}, {2, 4, 99}, {10, 4, 63}}
	queries := [][]int{{9, 7, 65}, {9, 6, 1}, {4, 5, 34}, {10, 8, 43}, {3, 7, 76}, {4, 2, 15}, {7, 6, 52}, {2, 0, 50}, {7, 6, 62}, {1, 0, 81}, {4, 5, 35}, {0, 11, 86}, {12, 5, 50}, {11, 2, 2}, {9, 5, 6}, {12, 0, 95}, {10, 6, 9}, {9, 4, 73}, {6, 10, 48}, {12, 0, 91}, {9, 10, 58}, {9, 8, 73}, {2, 3, 44}, {7, 11, 83}, {5, 3, 14}, {6, 2, 33}}
	fmt.Println(distanceLimitedPathsExist(13, edgeList, queries))
}

type DS struct {
	parent []int
}

func NewDS(n int) *DS {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	return &DS{
		parent: parent,
	}
}

func (s *DS) FindRoot(node int) int {
	// for s.parent[node] != -1 {
	//     node = s.parent[node]
	// }
	if s.parent[node] != -1 {
		s.parent[node] = s.FindRoot(s.parent[node])
	}
	if s.parent[node] == -1 {
		return node
	}
	return s.parent[node]
}

func (s *DS) Union(x int, y int) {
	xParent := s.FindRoot(x)
	yParent := s.FindRoot(y)
	if xParent == yParent {
		return
	}
	s.parent[yParent] = xParent
}

func (s *DS) IsConnected(x int, y int) bool {
	return s.FindRoot(x) == s.FindRoot(y)
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	lenEdge := len(edgeList)
	lenQuery := len(queries)

	for i := range queries {
		queries[i] = append(queries[i], i)
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	ds := NewDS(n)

	result := make([]bool, lenQuery)
	i := 0
	for _, q := range queries {
		for i < lenEdge && edgeList[i][2] < q[2] {
			ds.Union(edgeList[i][0], edgeList[i][1])
			i++
		}
		result[q[3]] = ds.IsConnected(q[0], q[1])
	}
	return result
}

// todo 看下索引排序为啥有问题
// func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
//     lenEdge := len(edgeList)
//     lenQuery := len(queries)
//     idxes := make([]int, lenQuery)
//     for i := range queries {
//         idxes[i] = i
//     }
//
//     sort.Slice(idxes, func(i, j int) bool {
//         return queries[i][2] < queries[j][2]
//     })
//     sort.Slice(edgeList, func(i, j int) bool {
//         return edgeList[i][2] < edgeList[j][2]
//     })
//
//
//     ds := NewDS(n)
//
//     result := make([]bool, lenQuery)
//     i := 0
//     for _, idx := range idxes {
//         for i < lenEdge && edgeList[i][2] < queries[idx][2] {
//             ds.Union(edgeList[i][0], edgeList[i][1])
//             i++
//         }
//         fmt.Println(ds.parent)
//         result[idx] = ds.IsConnected(queries[idx][0], queries[idx][1])
//     }
//     return result
// }
