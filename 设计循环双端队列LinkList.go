package main

func main() {

}

type Node struct {
	pre  *Node
	next *Node
	val  int
}

type MyCircularDeque struct {
	head *Node
	tail *Node
	size int
	len  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{size: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.len == this.size {
		return false
	}
	node := &Node{val: value, next: this.head}
	if this.len == 0 {
		this.tail = node
	} else {
		this.head.pre = node
	}
	this.head = node
	this.len++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.len == this.size {
		return false
	}
	node := &Node{val: value, pre: this.tail}
	if this.len == 0 {
		this.head = node
	} else {
		this.tail.next = node
	}
	this.tail = node
	this.len++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.len == 0 {
		return false
	}
	this.head = this.head.next
	this.len--
	if this.len == 0 {
		this.head = nil
		this.tail = nil
	} else {
		this.head.pre = nil
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.len == 0 {
		return false
	}
	this.tail = this.tail.pre
	this.len--
	if this.len == 0 {
		this.head = nil
		this.tail = nil
	} else {
		this.tail.next = nil
	}
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.len == 0 {
		return -1
	}
	return this.head.val
}

func (this *MyCircularDeque) GetRear() int {
	if this.len == 0 {
		return -1
	}
	return this.tail.val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.size
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
