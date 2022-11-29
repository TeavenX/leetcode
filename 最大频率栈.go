package main

func main() {

}

type FreqStack struct {
	freq    map[int]int   // 统计每一个key的频率
	group   map[int][]int // 将每个key根据不同的频率分组
	maxFreq int           // 记录最大的频率
}

func Constructor() FreqStack {
	return FreqStack{
		freq:    make(map[int]int),
		group:   make(map[int][]int),
		maxFreq: 0,
	}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	freq := this.freq[val]
	// 这里可以直接append的原因是pop的时候，只看maxFreq对应的
	// 所以lowerFreq存对应的值也没有关系，并且还能保留push的顺序
	this.group[freq] = append(this.group[freq], val)
	if this.maxFreq < freq {
		this.maxFreq = freq
	}
}

func (this *FreqStack) Pop() int {
	freq := this.maxFreq
	l := len(this.group[freq])
	val := this.group[freq][l-1]
	this.freq[val]--
	this.group[freq] = this.group[freq][:l-1]
	if l == 1 {
		this.maxFreq--
	}
	return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
