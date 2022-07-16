package main

func main() {

}

type MovingAverage struct {
	size      int
	windowSum int
	window    []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{size, 0, make([]int, 0)}
}

func (this *MovingAverage) Next(val int) float64 {
	this.windowSum += val
	this.window = append(this.window, val)
	if len(this.window)-1 < this.size {
		return float64(this.windowSum) / float64(len(this.window))
	} else {
		last := this.window[0]
		this.window = this.window[1:]
		this.windowSum -= last
		return float64(this.windowSum) / float64(this.size)
	}
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
