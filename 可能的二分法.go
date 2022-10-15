package main

func main() {

}

type GNode struct {
	Val  int
	Next []*GNode
	// 0: 未分配，-1: 组1，1: 组2
	color int
}

func possibleBipartition(n int, dislikes [][]int) bool {
	cache := make(map[int]*GNode)
	for i := 1; i <= n; i++ {
		cache[i] = &GNode{Val: i}
	}
	for _, tuple := range dislikes {
		cache[tuple[0]].Next = append(cache[tuple[0]].Next, cache[tuple[1]])
		cache[tuple[1]].Next = append(cache[tuple[1]].Next, cache[tuple[0]])
	}
	var dfs func(node *GNode, color int) bool
	dfs = func(node *GNode, color int) bool {
		node.color = color
		for _, next := range node.Next {
			if next.color == color {
				return false
			}
			if next.color == 0 && !dfs(next, -color) {
				return false
			}
		}
		return true
	}
	for i := 1; i <= n; i++ {
		if cache[i].color == 0 && !dfs(cache[i], 1) {
			return false
		}
	}
	return true
}
