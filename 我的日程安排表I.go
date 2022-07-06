package main

import "fmt"

func main() {
	c := Constructor()
	fmt.Println(c.Book(10, 20))
	fmt.Println(c)
	fmt.Println(c.Book(15, 25))
	fmt.Println(c)
	fmt.Println(c.Book(20, 30))

	//fmt.Println(c.Book(25, 30))
	//fmt.Println(c.Book(0, 5))
}

//type seg struct {
//	start, end  int
//	left, right *seg
//}
//
//type MyCalendar struct {
//	root *seg
//}
//
//func Constructor() MyCalendar {
//	return MyCalendar{root: &seg{start: -1, end: 0}}
//}
//
//func (c *MyCalendar) Book(start, end int) bool {
//	return c.insert(c.root, &seg{start: start, end: end})
//}
//
//func (c *MyCalendar) insert(cur, nw *seg) bool {
//	if nw.start < cur.end && nw.end > cur.start {
//		return false
//	}
//	if nw.end <= cur.start {
//		if cur.left == nil {
//			cur.left = nw
//			return true
//		}
//		return c.insert(cur.left, nw)
//	}
//	if cur.right == nil {
//		cur.right = nw
//		return true
//	}
//	return c.insert(cur.right, nw)
//}

type seg struct {
	val, lazy bool
}

type MyCalendar map[int]seg

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start, end int) bool {
	return c.insert(start, end-1, 0, 1e9, 1)
}

func (c MyCalendar) insert(start, end, left, right, idx int) bool {
	if start > right || end < left {
		return true
	}
	p := c[idx]
	if start <= left && end >= right {
		// 完全在当前区间内
		if p.val {
			return false
		}
		p.val = true
		c[idx] = p
		return true
	}
	// 有一部分在别的区间
	mid := (right-left)>>1 + left
	r1 := c.insert(start, end, left, mid, idx*2)
	r2 := c.insert(start, end, mid+1, right, idx*2+1)
	if r1 && r2 {
		p.val = true
		c[idx] = p
		return true
	}
	return false
}
