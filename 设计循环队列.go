package main

func main() {

}

type Node struct {
	Val  int
	Next *Node
}

type MyCircularQueue struct {
	len  int
	size int
	head *Node
	end  *Node
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{size: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.len == this.size {
		return false
	}
	node := &Node{Val: value}
	if this.len == 0 {
		this.head = node
	} else {
		end := this.end
		end.Next = node
	}
	this.end = node
	this.len++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.len == 0 {
		return false
	}
	this.len--
	if this.len == 0 {
		this.head = nil
		this.end = nil
	} else {
		this.head = this.head.Next
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.len == 0 {
		return -1
	}
	return this.head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.len == 0 {
		return -1
	}
	return this.end.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.len == this.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
