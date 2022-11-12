package main

func main() {

}

type node struct {
	x, y, k, d int
}

/*

n, m

queue := []node{}
queue = append(queue, node{x0, y0, 0, 0})

minD := math.MaxInt

for len(queue) > 0 {
    size := len(queue)
    for i := 0; i < size; i++ {
        nd := queue[i]
        if nd.k == targetK {
            minD = min(minD, nd.d)
        }
        for _, step := range steps {
            x, y := nd.x + step[0], nd.y + step[1]
            if x, y在格子内 && grid[x][y] != '#' {
                if grid[x][y] in 'A-F' && nd.k >> grid[x][y] - 'A' & 1 == 0 {
                    continue
                }
                nk := nd.k
                if grid[x][y] in 'a-f' {
                    nk |= 1 << grid[x][y]-'a'
                }
                visit[x][y][k] = true
                queue = append(queue, node{x, y, nk, d+1})
            }
        }
    }
    queue = queue[size:]
}
*/

var steps = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type node struct {
	x, y, k, d int
}

type tuple struct {
	x, y, k int
}

func shortestPathAllKeys(grid []string) int {
	start := node{}
	n, m := len(grid), len(grid[0])
	kc := 0
	visit := make(map[tuple]bool)
	for x, r := range grid {
		for y, s := range r {
			if s == '@' {
				start.x = x
				start.y = y
			}
			if 'a' <= s && s <= 'f' {
				kc |= 1 << (s - 'a')
			}
		}
	}
	queue := []node{start}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd.k == kc {
				return nd.d
			}
			for _, step := range steps {
				x, y := nd.x+step[0], nd.y+step[1]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] != '#' && !visit[tuple{x, y, nd.k}] {
					pos := grid[x][y]
					if 'A' <= pos && pos <= 'F' && nd.k&(1<<(pos-'A')) == 0 {
						continue
					}
					nk := nd.k
					if 'a' <= pos && pos <= 'f' {
						nk |= 1 << (pos - 'a')
					}
					visit[tuple{x, y, nk}] = true
					queue = append(queue, node{x, y, nk, nd.d + 1})
				}
			}
		}
		queue = queue[size:]
	}
	return -1
}
