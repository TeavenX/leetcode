package main

func main() {

}

type DisjointSet map[int]int

func (ds *DisjointSet) findRoot(node int) int {
	if p, ok := (*ds)[node]; ok {
		(*ds)[node] = ds.findRoot(p)
		return (*ds)[node]
	}
	return node
}

func (ds *DisjointSet) union(x, y int) {
	xp := ds.findRoot(x)
	yp := ds.findRoot(y)
	if xp != yp {
		(*ds)[yp] = xp
	}
}

func (ds *DisjointSet) isConnected(x, y int) bool {
	return ds.findRoot(x) == ds.findRoot(y)
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	ds := make(DisjointSet, n)
	for _, edge := range edges {
		ds.union(edge[0], edge[1])
	}
	return ds.isConnected(source, destination)
}
