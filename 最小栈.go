package main

import "fmt"

func main() {

	/*["MinStack","push","push","push","push","getMin","pop","getMin","pop","getMin","pop","getMin"]
	  [[],[2],[0],[3],[0],[],[],[],[],[],[],[]]*/

	//s := Constructor()
	//s.Push(2)
	//s.Push(0)
	//s.Push(3)
	//s.Push(0)
	//fmt.Println(s.GetMin())
	//s.Pop()
	//fmt.Println(s.GetMin())
	//s.Pop()
	//fmt.Println(s.GetMin())
	//s.Pop()
	//fmt.Println(s.GetMin())

	// ["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	// [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]

	s := Constructor()
	s.Push(2147483646)
	s.Push(2147483646)
	s.Push(2147483647)
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.GetMin())
	s.Pop()
	s.Push(2147483647)
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
	s.Push(-2147483648)
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.GetMin())
}

type MinStack struct {
	//min    int
	//queue  []int
	//minStk []int

	min   int64
	queue []int64
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	//if this.min > val {
	//	this.min = val
	//	this.minStk = append(this.minStk, val)
	//} else {
	//	this.minStk = append(this.minStk, this.min)
	//}
	//this.queue = append(this.queue, val)

	nVal := int64(val)
	if len(this.queue) == 0 {
		this.min = nVal
	}
	this.queue = append(this.queue, nVal-this.min)
	if this.min > nVal {
		this.min = nVal
	}
}

func (this *MinStack) Pop() {
	//this.minStk = this.minStk[:len(this.minStk)-1]
	//if len(this.minStk) == 0 {
	//	this.min = math.MaxInt
	//} else {
	//	this.min = this.minStk[len(this.minStk)-1]
	//}
	//this.queue = this.queue[:len(this.queue)-1]

	val := this.queue[len(this.queue)-1]
	if val < 0 {
		this.min = this.min - val
	}
	this.queue = this.queue[:len(this.queue)-1]
}

func (this *MinStack) Top() int {
	if this.queue[len(this.queue)-1] < 0 {
		return int(this.min)
	}
	temp := this.queue[len(this.queue)-1] + this.min
	return int(temp)
}

func (this *MinStack) GetMin() int {
	return int(this.min)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
