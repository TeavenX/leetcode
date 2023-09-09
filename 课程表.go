package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}
	var cache []bool
	var dfs func(target, i int) bool
	dfs = func(target, i int) bool {
		if i == target {
			return true
		}
		if cache[i] {
			return false
		}
		cache[i] = true
		for _, nxt := range g[i] {
			if dfs(target, nxt) {
				return true
			}
		}
		return false
	}
	for i := 0; i < numCourses; i++ {
		cache = make([]bool, numCourses)
		for _, nxt := range g[i] {
			if dfs(i, nxt) {
				return false
			}
		}
	}
	return true
}

func canFinishV2(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
		indeg[p[0]]++
	}
	q := make([]int, 0)
	for i, x := range indeg {
		if x == 0 {
			q = append(q, i)
		}
	}
	count := 0
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		count++
		for _, j := range g[i] {
			indeg[j]--
			if indeg[j] == 0 {
				q = append(q, j)
			}
		}
	}
	return count == numCourses
}
