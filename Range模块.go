package main

import "fmt"

func main() {
	rm := Constructor()
	//rm.AddRange(10, 20)
	//rm.RemoveRange(14, 16)
	//rm.QueryRange(10, 14)
	//rm.QueryRange(13, 15)
	//rm.QueryRange(16, 17)

	rm.AddRange(6, 8)
	rm.RemoveRange(7, 8)
	rm.RemoveRange(8, 9)
	rm.AddRange(8, 9)
	rm.RemoveRange(1, 3)
	rm.AddRange(1, 8)
	fmt.Println(rm.QueryRange(2, 4))
}

const N = 1e9 + 10

type Node struct {
	left, right *Node
	cover       bool
	add         bool
}

type RangeModule struct {
	root *Node
}

func Constructor() RangeModule {
	return RangeModule{
		root: &Node{nil, nil, false, false},
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	this.update(this.root, 1, N, left, right-1, 1)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.query(this.root, 1, N, left, right-1)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.update(this.root, 1, N, left, right-1, -1)
}

func (this *RangeModule) update(node *Node, start int, end int, left int, right int, val int) {
	if left <= start && end <= right {
		node.cover = 1 == val
		node.add = true
		return
	}
	mid := (start + end) >> 1
	this.pushDown(node)
	if left <= mid {
		this.update(node.left, start, mid, left, right, val)
	}
	if right > mid {
		this.update(node.right, mid+1, end, left, right, val)
	}
	this.pushUp(node)
}

func (this *RangeModule) query(node *Node, start int, end int, left int, right int) bool {
	if left <= start && end <= right {
		return node.cover
	}
	mid := (start + end) >> 1
	this.pushDown(node)
	result := true
	if left <= mid {
		result = result && this.query(node.left, start, mid, left, right)
	}
	if right > mid {
		result = result && this.query(node.right, mid+1, end, left, right)
	}
	return result
}

func (this *RangeModule) pushDown(node *Node) {
	if node.left == nil {
		node.left = &Node{}
	}
	if node.right == nil {
		node.right = &Node{}
	}
	if !node.add {
		return
	}
	node.left.cover, node.left.add, node.right.cover, node.right.add = node.cover, true, node.cover, true
	node.add = false
}

func (this *RangeModule) pushUp(node *Node) {
	node.cover = node.left.cover && node.right.cover
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
