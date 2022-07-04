package main

import "fmt"

func main() {
	c := Constructor()
	fmt.Println(c.Book(10, 20))
	fmt.Println(c)
	fmt.Println(c.Book(30, 35))
	fmt.Println(c)
	fmt.Println(c.Book(20, 35))
	fmt.Println(c)
	fmt.Println(c.Book(35, 50))
	fmt.Println(c)
}

type pair struct {
	num, lazy int
}

type MyCalendarThree map[int]pair

func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}

func (this MyCalendarThree) update(start, end, left, right, idx int) {
	if left > end || right < start {
		return
	}
	if start <= left && right <= end {
		p := this[idx]
		p.num++
		p.lazy++
		this[idx] = p
	} else {
		mid := (left + right) >> 1
		// 左子树
		this.update(start, end, left, mid, idx*2)
		// 右子树
		this.update(start, end, mid+1, right, idx*2+1)
		p := this[idx]
		p.num = p.lazy + max(this[idx*2].num, this[idx*2+1].num)
		this[idx] = p
	}
}

func (this MyCalendarThree) Book(start int, end int) int {
	this.update(start, end-1, 0, 1e9, 1)
	return this[1].num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
