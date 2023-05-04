package main

func main() {

}

func hardestWorker(n int, logs [][]int) int {
	id := logs[0][0]
	t := logs[0][1]
	for i, l := range logs[1:] {
		if v := l[1] - logs[i][1]; v >= t {
			if v == t {
				id = min(id, l[0])
			} else {
				id = l[0]
			}
			t = v
		}
	}
	return id
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
