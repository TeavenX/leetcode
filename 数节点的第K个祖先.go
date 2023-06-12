package main

import "math/bits"

func main() {

}

// TLE
// type TreeAncestor struct {
// 	data  []int
// 	cache map[int]int
// }
//
// func Constructor(n int, parent []int) TreeAncestor {
// 	return TreeAncestor{
// 		data:  parent,
// 		cache: make(map[int]int),
// 	}
// }
//
// func (this *TreeAncestor) GetKthAncestor(node int, k int) (res int) {
// 	key := node*1e6 + k
// 	if v, ok := this.cache[key]; ok {
// 		return v
// 	}
// 	defer func() {
// 		this.cache[key] = res
// 	}()
// 	p := this.data[node]
// 	k--
// 	for k > 0 && p > -1 {
// 		k--
// 		p = this.data[p]
// 	}
// 	return p
// }

type TreeAncestor [][]int

func Constructor(n int, parent []int) TreeAncestor {
	m := bits.Len(uint(n))
	data := make([][]int, n)
	for i, p := range parent {
		data[i] = make([]int, m)
		data[i][0] = p
	}
	for i := 0; i < m-1; i++ { // 指数
		for x := 0; x < n; x++ { // node
			if p := data[x][i]; p != -1 {
				data[x][i+1] = data[p][i]
			} else {
				data[x][i+1] = -1
			}
		}
	}
	return data
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	data := *this
	m := bits.Len(uint(k))
	for i := 0; i < m; i++ {
		if k>>i&1 == 1 {
			node = data[node][i]
			if node == -1 {
				break
			}
		}
	}
	return node
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
