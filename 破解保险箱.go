package main

import (
	"bytes"
	"math"
	"strconv"
)

func main() {

}

func crackSafe(n int, k int) string {
	// edgeNum := int(math.Pow(float64(k), float64(n)))
	nodeNum := int(math.Pow(float64(k), float64(n-1)))
	edges := make([]int, nodeNum)
	for i := 0; i < nodeNum; i++ {
		edges[i] = k - 1
	}
	buf := bytes.NewBufferString("")
	for i := 0; i < n-1; i++ {
		buf.WriteByte('0')
	}
	node := 0
	for edges[node] >= 0 {
		edge := edges[node]
		edges[node]--
		node = (node*k + edge) % nodeNum
		buf.WriteString(strconv.Itoa(edge))
	}
	return buf.String()
}
