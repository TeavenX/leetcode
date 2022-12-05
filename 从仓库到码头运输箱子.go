package main

func main() {

}

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	p := make([]int, n+1)   // 目的地的分布
	w := make([]int, n+1)   // 重量的分布
	W := make([]int, n+1)   // 重量的前缀和
	neg := make([]int, n+1) // 前后目的地不一致前缀和

	for i := 1; i <= n; i++ {
		p[i] = boxes[i-1][0]
		w[i] = boxes[i-1][1]
		if i > 1 {
			neg[i] = neg[i-1]
			if p[i] != p[i-1] {
				neg[i]++
			}
		}
		W[i] = W[i-1] + w[i]
	}

	deque := []int{0}
	f := make([]int, n+1)
	g := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for len(deque) > 0 && (i-deque[0] > maxBoxes || W[i]-W[deque[0]] > maxWeight) {
			deque = deque[1:]
		}

		f[i] = g[deque[0]] + neg[i] + 2

		if i != n {
			g[i] = f[i] - neg[i+1]
			for len(deque) > 0 && g[i] <= g[deque[len(deque)-1]] {
				deque = deque[:len(deque)-1]
			}
			deque = append(deque, i)
		}
	}
	return f[n]
}
