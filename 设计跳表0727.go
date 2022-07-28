package main

import "math/rand"

func main() {

}

type Node struct {
	Val  int
	Next *Node
	Down *Node
}

type Skiplist struct {
	Head *Node
}

func Constructor() Skiplist {
	return Skiplist{Head: &Node{Val: -1}}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.Head
	for cur != nil {
		for cur.Next != nil && cur.Next.Val < target {
			cur = cur.Next
		}
		if cur.Next != nil && cur.Next.Val == target {
			return true
		}
		cur = cur.Down
	}
	return false
}

func (this *Skiplist) Add(num int) {
	cur, insert := this.Head, true
	down := &Node{Val: -1}
	queue := make([]*Node, 0)
	for cur != nil {
		for cur.Next != nil && cur.Next.Val < num {
			cur = cur.Next
		}
		queue = append(queue, cur)
		cur = cur.Down
	}
	for insert && len(queue) > 0 {
		cur = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if down.Val == -1 {
			cur.Next = &Node{Val: num, Next: cur.Next}
		} else {
			cur.Next = &Node{Val: num, Next: cur.Next, Down: down}
		}
		if rand.Float64() < 0.5 {
			insert = false
		}
		down = cur.Next
	}
	if insert {
		this.Head = &Node{Val: -1, Down: this.Head}
	}
}

func (this *Skiplist) Erase(num int) bool {
	cur, found := this.Head, false
	for cur != nil {
		for cur.Next != nil && cur.Next.Val < num {
			cur = cur.Next
		}
		if cur.Next != nil && cur.Next.Val == num {
			found = true
			cur.Next = cur.Next.Next
		}
		cur = cur.Down
	}
	return found
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
