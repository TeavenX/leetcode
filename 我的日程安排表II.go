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

//type pair struct {
//	num, lazy int
//}

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

type SegmentTreeNode struct {
	left, right *SegmentTreeNode
	val, lazy   int
}

type MyCalendarTwo struct {
	root SegmentTreeNode
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{root: SegmentTreeNode{}}
}

func (this *SegmentTreeNode) update(start, end, left, right int) {
	if start <= left && right <= end {
		this.val++
		this.lazy++
		return
	}
	this.pushdown()
	mid := (right-left)>>1 + left
	if mid >= start {
		this.left.update(start, end, left, mid)
	}
	if mid < end {
		this.right.update(start, end, mid+1, right)
	}
	this.pushup()
}

func (this *SegmentTreeNode) pushdown() {
	if this.left == nil {
		this.left = &SegmentTreeNode{}
	}
	if this.right == nil {
		this.right = &SegmentTreeNode{}
	}
	if this.lazy > 0 {
		this.left.val += this.lazy
		this.left.lazy += this.lazy
		this.right.val += this.lazy
		this.right.lazy += this.lazy
		this.lazy = 0
	}
}

func (this *SegmentTreeNode) pushup() {
	this.val = max(this.left.val, this.right.val)
}

func (this *SegmentTreeNode) query(start, end, left, right int) int {
	if start <= left && end >= right {
		return this.val
	}
	this.pushdown()
	mid := (right-left)>>1 + left
	result := 0
	if start <= mid {
		result = this.left.query(start, end, left, mid)
	}
	if end > mid {
		result = max(result, this.right.query(start, end, mid+1, right))
	}
	return result
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if this.root.query(start, end-1, 0, 1e9) >= 2 {
		return false
	}
	this.root.update(start, end-1, 0, 1e9)
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
