package main

import (
	"container/heap"
	"math"
)

func main() {

}

type Pair struct {
	b, v, i, count int
}

type HP []Pair

func (h HP) Len() int {
	return len(h)
}

func (h HP) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h HP) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HP) Push(v interface{}) {
	*h = append(*h, v.(Pair))
}

func (h *HP) Pop() interface{} {
	t := *h
	v := t[len(t)-1]
	t = t[:len(t)-1]
	*h = t
	return v
}

func storeWaterError(bucket []int, vat []int) int {
	n := len(bucket)
	base := 0
	hp := HP(make([]Pair, 0, n))
	for i := range bucket {
		if vat[i] == 0 {
			continue
		}
		if bucket[i] == 0 {
			bucket[i] = 1
			base++
		}
		count := (vat[i] + bucket[i] - 1) / bucket[i]
		heap.Push(&hp, Pair{bucket[i], vat[i], i, count})
	}
	ans := math.MaxInt
	for {
		p := heap.Pop(&hp).(Pair)
		p.b++
		p.count = (p.v + p.b - 1) / p.b
		res := base + p.count + p.b - bucket[p.i]
		if res <= ans {
			ans = res
		} else {
			return ans
		}
		heap.Push(&hp, p)
	}
}

func storeWater(bucket []int, vat []int) int {
	mx := 0
	for _, v := range vat {
		mx = max(mx, v)
	}
	if mx == 0 {
		return 0
	}
	ans := math.MaxInt
	for c := 1; c <= mx; c++ {
		// c 枚举倒的次数
		res := 0
		for i, v := range vat {
			// 计算倒c次，需要多大的桶，减去当前桶容量就是步骤数
			// 累加桶容量变更的操作数
			res += max(0, (v+c-1)/c-bucket[i])
		}
		// 总操作数res + 倒水次数c就是结果
		ans = min(ans, res+c)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
