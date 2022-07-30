package main

func main() {
	//nums := []int{4, 6, 15, 35}
	nums := []int{801, 839, 521, 999, 655, 16, 536, 84, 670, 983, 408, 346, 511, 638, 895}
	largestComponentSize(nums)
}

type Node struct {
	Val    int
	Parent *Node
	Rank   int
}

func (n *Node) findRoot() *Node {
	node := n
	for node.Parent != nil {
		node = node.Parent
	}
	return node
}

type disjointSet map[int]*Node

func (d disjointSet) union(a, b int) {
	pa := d[a].findRoot()
	pb := d[b].findRoot()
	if pa == pb {
		return
	}
	if pa.Rank < pb.Rank {
		pa.Parent = pb
		d[a].Parent = pb
	} else if pa.Rank > pb.Rank {
		pb.Parent = pa
		d[b].Parent = pa
	} else {
		pb.Parent = pa
		d[b].Parent = pa
		pa.Rank++
	}
}

func largestComponentSize(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	ds := disjointSet{}
	for _, num := range nums {
		ds[num] = &Node{Val: num, Rank: 1}
	}

	for i, x := range nums {
		for _, y := range nums[i+1:] {
			if calCommonFactor(x, y) {
				ds.union(x, y)
			}
		}
	}
	result := 0
	cache := map[int]int{}
	for _, num := range nums {
		p := ds[num].findRoot().Val
		cache[p]++
		if cache[p] > result {
			result = cache[p]
		}
	}
	return result
}

func calCommonFactor(a, b int) bool {
	if a > b {
		a, b = b, a
	}
	for i := 2; i <= a; i++ {
		if a%i == 0 && b%i == 0 {
			return true
		}
	}
	return false
}
