package main

import "container/list"

func main() {

}

type Pair struct {
	k, v int
}

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	data  *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*list.Element, capacity),
		data:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	c, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.data.MoveToFront(c)
	r := c.Value.(*Pair)
	return r.v
}

func (this *LRUCache) Put(key int, value int) {
	c, ok := this.cache[key]
	if !ok {
		p := &Pair{key, value}
		this.data.PushFront(p)
		this.cache[key] = this.data.Front()
		if this.data.Len() > this.cap {
			back := this.data.Back()
			delete(this.cache, back.Value.(*Pair).k)
			this.data.Remove(back)
		}
		return
	}
	r := c.Value.(*Pair)
	r.v = value
	this.data.MoveToFront(c)
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
