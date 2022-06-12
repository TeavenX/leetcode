package main

import "fmt"

func main() {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	//words := []string{"abc", "def"}
	fmt.Println(alienOrderPlatForm(words))
	fmt.Println(alienOrderPlatFormBFS(words))
	fmt.Println(alienOrder(words))
}

func alienOrderPlatForm(words []string) string {
	g := map[byte][]byte{}
	for _, c := range words[0] {
		g[byte(c)] = nil
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, c := range t {
			// 把每一个字母都在字典中初始化，如果字典中已经存在则直接延用
			g[byte(c)] = g[byte(c)]
		}
		for j := 0; j < len(s) && j < len(t); j++ {
			if s[j] != t[j] {
				g[s[j]] = append(g[s[j]], t[j])
				continue next
			}
		}
		// 如果s和t前缀相同，且前一个单词的长度比后一个单词的长度长，不符合规则，直接返回
		if len(s) > len(t) {
			return ""
		}
	}

	const visiting, visited = 1, 2
	order := make([]byte, len(g))
	i := len(g) - 1
	state := map[byte]int{}
	var dfs func(u byte) bool
	dfs = func(u byte) bool {
		state[u] = visiting
		for _, v := range g[u] {
			if state[v] == 0 {
				if !dfs(v) {
					return false
				}
			} else if state[v] == visiting {
				return false
			}
		}
		order[i] = u
		i--
		state[u] = visited
		return true
	}
	for u := range g {
		if state[u] == 0 && !dfs(u) {
			return ""
		}
	}
	return string(order)
}

func alienOrderPlatFormBFS(words []string) string {
	// 声明邻接表
	g := map[byte][]byte{}
	// 记录每个节点的入度（有向图中指向该节点的边数量）
	inDeg := map[byte]int{}
	for _, c := range words[0] {
		inDeg[byte(c)] = 0
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, c := range t {
			inDeg[byte(c)] += 0
		}
		for j := 0; j < len(s) && j < len(t); j++ {
			if s[j] != t[j] {
				g[s[j]] = append(g[s[j]], t[j])
				inDeg[t[j]]++
				continue next
			}
		}
		if len(s) > len(t) {
			return ""
		}
	}

	order := make([]byte, len(inDeg))
	q := order[:0]
	for u, d := range inDeg {
		// 将入度为0的节点放入队列
		if d == 0 {
			q = append(q, u)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			// 将该节点所有相邻的节点的入度减一
			// 判断减去1之后，相邻节点的入度是否为0
			// 如果为0，则放入队列中
			// 如果队列遍历结束，邻接表的中还有入度非0的节点存在
			// 则证明该图不是有向无环图（存在环形依赖），不满足条件
			if inDeg[v]--; inDeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if cap(q) == 0 {
		return string(order)
	}
	return ""
}

func alienOrder(words []string) string {
	n := len(words)
	g := map[byte][]byte{}
	inDeg := map[byte]int{}

	for _, str := range words[0] {
		inDeg[byte(str)] = 0
	}

nextIter:
	for i := 1; i < n; i++ {
		pre, cur := words[i-1], words[i]
		for _, str := range cur {
			inDeg[byte(str)] += 0
		}
		for j := 0; j < len(pre) && j < len(cur); j++ {
			if pre[j] != cur[j] {
				g[pre[j]] = append(g[pre[j]], cur[j])
				inDeg[cur[j]]++
				continue nextIter
			}
		}
		if len(pre) > len(cur) {
			return ""
		}
	}
	result := make([]byte, len(inDeg))
	queue := result[:0]
	for node, d := range inDeg {
		if d == 0 {
			queue = append(queue, node)
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, cn := range g[node] {
			inDeg[cn]--
			if inDeg[cn] == 0 {
				queue = append(queue, cn)
			}
		}
	}
	if cap(queue) == 0 {
		return string(result)
	}
	return ""
}
