package main

func main() {

}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		a, b := p[0], p[1]
		g[b] = append(g[b], a)
		indeg[a]++
	}
	ans := make([]int, 0, numCourses)
	q := make([]int, 0)
	for i, x := range indeg {
		if x == 0 {
			q = append(q, i)
			ans = append(ans, i)
		}
	}
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		for _, j := range g[i] {
			indeg[j]--
			if indeg[j] == 0 {
				q = append(q, j)
				ans = append(ans, j)
			}
		}
	}
	if len(ans) != numCourses {
		return []int{}
	}
	return ans
}
