package main

import "fmt"

func main() {
	points := [][]int{{0, 1}, {0, 0}}
	points = [][]int{{4, 5}, {4, -1}, {4, 0}}
	fmt.Println(maxPoints(points))
}

func maxPointsError(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	cache := make([][]int, 0)
	xCache := make(map[int]int)
	max := 0
	maxS := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			x0, y0 := points[i][0], points[i][1]
			x1, y1 := points[j][0], points[j][1]
			if x0-x1 != 0 {
				k := (y0 - y1) / (x0 - x1)
				b := y0 - k*x0
				cache = append(cache, []int{k, b, 0})
			} else {
				xCache[x0]++
				if maxS < xCache[x0] {
					maxS = xCache[x0]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		for j := 0; j < len(cache); j++ {
			k, b := cache[j][0], cache[j][1]
			if y == k*x+b {
				cache[j][2]++
				if max < cache[j][2] {
					max = cache[j][2]
				}
			}
		}
	}
	if maxS == 1 {
		maxS = 2
	}
	if maxS > max {
		max = maxS
	}
	return max
}

func maxPoints(points [][]int) (ans int) {
	n := len(points)
	if n <= 2 {
		return n
	}
	for i, p := range points {
		if ans >= n-i || ans > n/2 {
			break
		}
		cnt := map[int]int{}
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			cnt[y+x*20001]++
		}
		for _, c := range cnt {
			ans = max(ans, c+1)
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPoints(points [][]int) int {
	n := len(points)
	ans := 1
	for i := 0; i < n; i++ {
		x := points[i]
		for j := i + 1; j < n; j++ {
			y := points[j]
			cnt := 2
			for k := j + 1; k < n; k++ {
				z := points[k]
				s1 := (x[0] - y[0]) * (y[1] - z[1])
				s2 := (x[1] - y[1]) * (y[0] - z[0])
				if s1 == s2 {
					cnt++
				}
			}
			ans = Max(ans, cnt)
		}
	}
	return ans
}
func Max(args ...int) int {
	max := args[0]
	for _, item := range args {
		if item > max {
			max = item
		}
	}
	return max
}
