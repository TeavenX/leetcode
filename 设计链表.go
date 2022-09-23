package main

func main() {
	l := Constructor()
	//l.AddAtHead(1)
	//l.AddAtTail(3)
	//l.AddAtIndex(1, 2)
	//l.Get(1)
	//l.DeleteAtIndex(1)
	//l.Get(1)

	//l.AddAtHead(1)
	//l.DeleteAtIndex(0)

	//l.AddAtHead(7)
	//l.AddAtHead(2)
	//l.AddAtHead(1)
	//l.AddAtIndex(3, 0)
	//l.DeleteAtIndex(2)
	//l.AddAtHead(6)
	//l.AddAtTail(4)
	//l.Get(4)
	//l.AddAtHead(4)
	//l.AddAtIndex(5, 0)
	//l.AddAtHead(6)

	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	l.Get(1)
	l.DeleteAtIndex(0)
	l.Get(0)
}

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Len  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Len {
		return -1
	}
	node := this.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val, Next: this.Head}
	this.Head = node
	this.Len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	if this.Len == 0 {
		this.Len++
		this.Head = node
		return
	}
	n := this.Head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = node
	this.Len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.Len {
		this.AddAtTail(val)
		return
	} else if index <= 0 {
		this.AddAtHead(val)
		return
	} else if index >= this.Len {
		return
	}
	node := &Node{Val: val}
	n := this.Head
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	node.Next = n.Next
	n.Next = node
	this.Len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Len {
		return
	}
	if index == 0 {
		this.Head = this.Head.Next
		this.Len--
		return
	}
	node := this.Head
	for i := 0; i < index-1; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	this.Len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
