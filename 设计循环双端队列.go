package main

func main() {

}

type Node struct {
}

type MyCircularDeque struct {
	data []int
	size int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{size: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.data) == this.size {
		return false
	}
	this.data = append([]int{value}, this.data...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.data) == this.size {
		return false
	}
	this.data = append(this.data, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[:len(this.data)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}

func (this *MyCircularDeque) GetRear() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return len(this.data) == this.size
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
