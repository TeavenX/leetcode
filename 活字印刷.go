package main

func main() {

}

func numTilePossibilities(tiles string) int {
	n := len(tiles)
	cache := make(map[string]bool)
	var dfs func(i int, pre string, used int)
	dfs = func(i int, pre string, used int) {
		if i == n {
			return
		}
		if cache[pre] {
			return
		}
		for j := 0; j < n; j++ {
			if used>>j&1 == 1 {
				continue
			}
			dfs(i+1, pre+string(tiles[j]), used|1<<j)
		}
		cache[pre] = true
	}
	dfs(-1, "", 0)
	return len(cache) - 1
}
