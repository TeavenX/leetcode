package main

func main() {

}

func maxDistToClosest(seats []int) int {
	pos := make([]int, 0)
	for i, seat := range seats {
		if seat == 1 {
			pos = append(pos, i)
		}
	}
	n, m := len(seats), len(pos)
	ans := max(pos[0], n-1-pos[len(pos)-1])
	for i := 0; i < m-1; i++ {
		ans = max(ans, (pos[i+1]-pos[i])>>1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
