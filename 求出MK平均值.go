package main

import "github.com/emirpasic/gods/trees/redblacktree"

func main() {

}

type MKAverage struct {
	q              []int
	low, mid, high *redblacktree.Tree
	s, m, k        int
	sizeL, sizeH   int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		low:  redblacktree.NewWithIntComparator(),
		mid:  redblacktree.NewWithIntComparator(),
		high: redblacktree.NewWithIntComparator(),
		m:    m,
		k:    k,
	}
}

func (this *MKAverage) AddElement(num int) {
	merge := func(rbt *redblacktree.Tree, key int, value int) {
		if v, ok := rbt.Get(key); ok {
			nv := v.(int) + value
			if nv == 0 {
				rbt.Remove(key)
			} else {
				rbt.Put(key, nv)
			}
		} else {
			rbt.Put(key, value)
		}
	}
	if this.low.Empty() || this.low.Right().Key.(int) >= num {
		merge(this.low, num, 1)
		this.sizeL++
	} else if this.high.Empty() || this.high.Left().Key.(int) <= num {
		merge(this.high, num, 1)
		this.sizeH++
	} else {
		merge(this.mid, num, 1)
		this.s += num
	}
	this.q = append(this.q, num)
	// 否则将数据流中最后 m 个元素拷贝到一个独立的容器中
	if len(this.q) > this.m {
		x := this.q[0]
		this.q = this.q[1:]
		if _, ok := this.low.Get(x); ok {
			merge(this.low, x, -1)
			this.sizeL--
		} else if _, ok = this.high.Get(x); ok {
			merge(this.high, x, -1)
			this.sizeH--
		} else {
			merge(this.mid, x, -1)
			this.s -= x
		}
	}

	// 这里改成if也可以
	// 从这个容器中删除最小的 k 个数和最大的 k 个数。
	for ; this.sizeL > this.k; this.sizeL-- {
		x := this.low.Right().Key.(int)
		merge(this.low, x, -1)
		merge(this.mid, x, 1)
		this.s += x
	}
	for ; this.sizeH > this.k; this.sizeH-- {
		x := this.high.Left().Key.(int)
		merge(this.high, x, -1)
		merge(this.mid, x, 1)
		this.s += x
	}
	// 如果不够就从mid中拿
	for ; this.sizeL < this.k && !this.mid.Empty(); this.sizeL++ {
		x := this.mid.Left().Key.(int)
		merge(this.low, x, 1)
		merge(this.mid, x, -1)
		this.s -= x
	}
	for ; this.sizeH < this.k && !this.mid.Empty(); this.sizeH++ {
		x := this.mid.Right().Key.(int)
		merge(this.high, x, 1)
		merge(this.mid, x, -1)
		this.s -= x
	}
}

func (this *MKAverage) CalculateMKAverage() int {
	// 如果数据流中的整数少于 m 个，MK 平均值 为 -1
	if len(this.q) < this.m {
		return -1
	}
	return this.s / (this.m - 2*this.k)
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
