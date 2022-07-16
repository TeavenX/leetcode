package main

import "fmt"

func main() {
	words := []string{"apple", "ape", "c", "c", "c"}
	w := Constructor(words)
	//fmt.Println(w.root)
	//fmt.Println(w.F("a", "e"))
	fmt.Println(w.F("c", "c"))
}

type Node struct {
	idx  int
	end  bool
	dict [26]*Node
}

func (n *Node) String() string {
	base := fmt.Sprintf("[idx: %d, end: %v, dict: \n\t %v]", n.idx, n.end, n.dict)
	return base
}

type WordFilter struct {
	row   []string
	root  *Node
	cache map[string]int
}

func Constructor(words []string) WordFilter {
	root := &Node{}
	for idx, word := range words {
		cur := root
		for _, s := range word {
			if cur.dict[s-'a'] == nil {
				cur.dict[s-'a'] = &Node{}
			}
			cur = cur.dict[s-'a']
		}
		cur.end = true
		cur.idx = idx
	}
	return WordFilter{root: root, row: words, cache: make(map[string]int)}
}

func (this *WordFilter) search(node *Node, pref, suff string) []int {
	for _, s := range pref {
		if node.dict[s-'a'] == nil {
			return []int{}
		}
		node = node.dict[s-'a']
	}
	result := make([]int, 0)
	var dfs func(cur *Node)
	dfs = func(cur *Node) {
		if cur == nil {
			return
		}
		if cur.end {
			w := this.row[cur.idx]
			if len(w) >= len(suff) && w[len(w)-len(suff):] == suff {
				result = append(result, cur.idx)
			}
		}
		for _, n := range cur.dict {
			dfs(n)
		}
	}
	dfs(node)
	return result
}

func (this *WordFilter) F(pref string, suff string) int {
	c := fmt.Sprintf("%s-%s", pref, suff)
	if result, ok := this.cache[c]; ok {
		return result
	}
	result := this.search(this.root, pref, suff)
	if len(result) == 0 {
		this.cache[c] = -1
		return -1
	}
	maxI := result[0]
	for _, i := range result {
		if i > maxI {
			maxI = i
		}
	}
	this.cache[c] = maxI
	return maxI
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
