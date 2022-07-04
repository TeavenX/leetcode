package main

import "fmt"

func main() {
	c := Constructor()
	fmt.Println(c.Book(10, 20))
	fmt.Println(c.Book(25, 30))
	fmt.Println(c.Book(20, 30))
	fmt.Println(c)
	fmt.Println(c.Book(0, 5))
}

type seg struct {
	start, end  int
	left, right *seg
}

type MyCalendar struct {
	root *seg
}

func Constructor() MyCalendar {
	return MyCalendar{root: &seg{start: -1, end: 0}}
}

func (c *MyCalendar) Book(start, end int) bool {
	return c.insert(c.root, &seg{start: start, end: end})
}

func (c *MyCalendar) insert(cur, nw *seg) bool {
	if nw.start < cur.end && nw.end > cur.start {
		return false
	}
	if nw.end <= cur.start {
		if cur.left == nil {
			cur.left = nw
			return true
		}
		return c.insert(cur.left, nw)
	}
	if cur.right == nil {
		cur.right = nw
		return true
	}
	return c.insert(cur.right, nw)
}
