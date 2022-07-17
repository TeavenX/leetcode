package main

func main() {

}

const (
	STATUSCLEAN = iota
	STATUSINFECT
	STATUSDONE
	STATUSINFECTVISITED
	STATUSCLEANVISITED
)

var steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func containVirusError(isInfected [][]int) int {
	n, m := len(isInfected), len(isInfected[0])
	wallCount := 0
	var dfs func(x, y int, source, target int) int
	for {
		visited := [55][55]int{}
		dfs = func(x, y int, source, target int) int {
			cleanCount := 0
			isInfected[x][y] = target
			for _, step := range steps {
				newx, newy := x+step[0], y+step[1]
				if newx >= 0 && newx < n && newy >= 0 && newy < m {
					if isInfected[newx][newy] == source && visited[newx][newy] != 1 {
						visited[newx][newy] = 1
						cleanCount += dfs(newx, newy, source, target)
					} else if isInfected[newx][newy] == STATUSCLEAN {
						cleanCount++
					}
				}
			}
			return cleanCount
		}
		maxCleanArea := 0
		curCleanArr := [2]int{}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if isInfected[i][j] == STATUSINFECT {
					cleanSize := dfs(i, j, STATUSINFECT, STATUSINFECT)
					if cleanSize > maxCleanArea {
						maxCleanArea = cleanSize
						curCleanArr[0], curCleanArr[1] = i, j
					}
				}
			}
		}
		if maxCleanArea == 0 {
			break
		}
		dfs(curCleanArr[0], curCleanArr[1], STATUSINFECT, STATUSDONE)
		wallCount += maxCleanArea
	}
	return wallCount
}

type pair struct {
	x, y int
}

func containVirusCV(isInfected [][]int) int {
	n, m := len(isInfected), len(isInfected[0])
	result := 0
	for {
		neighbors := make([]map[pair]struct{}, 0)
		firewalls := make([]int, 0)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if isInfected[i][j] != 1 {
					continue
				}
				q := []pair{{i, j}}
				neighbor := map[pair]struct{}{}
				firewall, idx := 0, len(neighbors)+1
				isInfected[i][j] = -idx
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					for _, step := range steps {
						newx, newy := p.x+step[0], p.y+step[1]
						if newx >= 0 && newy >= 0 && newx < n && newy < m {
							if isInfected[newx][newy] == 1 {
								q = append(q, pair{newx, newy})
								isInfected[newx][newy] = -idx
							} else if isInfected[newx][newy] == 0 {
								firewall++
								neighbor[pair{newx, newy}] = struct{}{}
							}
						}
					}
				}
				neighbors = append(neighbors, neighbor)
				firewalls = append(firewalls, firewall)
			}
		}
		if len(neighbors) == 0 {
			break
		}
		idx := 0
		for i := 1; i < len(neighbors); i++ {
			if len(neighbors[i]) > len(neighbors[idx]) {
				idx = i
			}
		}

		result += firewalls[idx]
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if isInfected[i][j] < 0 {
					if isInfected[i][j] != -idx-1 {
						isInfected[i][j] = 1
					} else {
						isInfected[i][j] = 2
					}
				}
			}
		}

		for i, neighbor := range neighbors {
			if i != idx {
				for p := range neighbor {
					isInfected[p.x][p.y] = 1
				}
			}
		}

		if len(neighbors) == 1 {
			break
		}
	}
	return result
}
