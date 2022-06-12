package main

func main() {

}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	result := make([][]int, 0)
	var dfs func(int)
	temp := make([]int, 0)
	dfs = func(idx int) {
		if idx == n {
			return
		}
		temp = append(temp, idx)
		if idx == n-1 {
			a := make([]int, len(temp))
			copy(a, temp)
			result = append(result, a)
		}
		for _, val := range graph[idx] {
			dfs(val)
		}
		temp = temp[:len(temp)-1]
	}
	dfs(0)
	return result
}
