package main

func main() {

}

type Pair struct {
	i, h, w int
}

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	cache := make(map[Pair]int)
	var dfs func(i int, height int, width int) int
	dfs = func(i int, height int, width int) (res int) {
		if i == n {
			return height
		}
		p := Pair{i, height, width}
		if v, ok := cache[p]; ok {
			return v
		}
		defer func() {
			cache[p] = res
		}()
		w, h := books[i][0], books[i][1]
		if w+width > shelfWidth {
			return height + dfs(i+1, h, w)
		}
		return min(dfs(i+1, max(height, h), width+w), height+dfs(i+1, h, w))
	}
	return dfs(0, 0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
