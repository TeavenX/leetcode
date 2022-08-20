package main

func main() {

}

type OrderedStream struct {
	data []string
	cur  int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{data: make([]string, n+1), cur: 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.data[idKey] = value
	cur := this.cur
	for this.cur < len(this.data) && this.data[this.cur] != "" {
		this.cur++
	}
	return this.data[cur:this.cur]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
