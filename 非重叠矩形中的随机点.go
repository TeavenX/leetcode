package main

import (
	"fmt"
	"math/rand"
)

func main() {
	test()
}

type Solution struct {
	rects [][]int
}

func Constructor(rects [][]int) Solution {
	return Solution{
		rects: rects,
	}
}

func (this *Solution) Pick() []int {
	total := 0
	choices := this.rects[0]
	for _, rect := range this.rects {
		area := (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
		total += area
		if rand.Intn(total) <= area {
			choices = rect
		}
	}
	x := rand.Intn(choices[2]-choices[0]+1) + choices[0]
	y := rand.Intn(choices[3]-choices[1]+1) + choices[1]
	return []int{x, y}
}

//type Solution struct {
//	rects [][]int
//	sum   []int
//}
//
//func Constructor(rects [][]int) Solution {
//	sum := make([]int, len(rects)+1)
//	for i, r := range rects {
//		a, b, x, y := r[0], r[1], r[2], r[3]
//		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
//	}
//	return Solution{rects, sum}
//}
//
//func (s *Solution) Pick() []int {
//	k := rand.Intn(s.sum[len(s.sum)-1])
//	rectIndex := sort.SearchInts(s.sum, k+1) - 1
//	r := s.rects[rectIndex]
//	a, b, y := r[0], r[1], r[3]
//	da := (k - s.sum[rectIndex]) / (y - b + 1)
//	db := (k - s.sum[rectIndex]) % (y - b + 1)
//	return []int{a + da, b + db}
//}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */

func test() {
	result := make(map[int]int)
	rects := [][]int{{82918473, -57180867, 82918476, -57180863}, {83793579, 18088559, 83793580, 18088560}, {66574245, 26243152, 66574246, 26243153}, {72983930, 11921716, 72983934, 11921720}}
	s := Constructor(rects)
	for i := 0; i < 1e4; i++ {
		tmp := s.Pick()
		flag := false
		for idx, rect := range rects {
			if tmp[0] >= rect[0] && tmp[0] <= rect[2] && tmp[1] >= rect[1] && tmp[1] <= rect[3] {
				result[idx]++
				flag = true
			}
		}
		if !flag {
			fmt.Println(tmp)
		}
	}
	fmt.Println(result)
}
