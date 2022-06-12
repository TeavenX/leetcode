package main

import "fmt"

func main() {
	rc := Constructor()
	list := []int{642, 1849, 4921, 5936, 5957}
	for _, val := range list {
		fmt.Println(rc.Ping(val))
	}
}

type RecentCounter struct {
	Cache []int
}

func Constructor() RecentCounter {
	//return RecentCounter{make([]int, 0, 3001)}
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	if len(this.Cache) == 0 {
		this.Cache = append(this.Cache, t)
	} else {
		left, right := 0, len(this.Cache)-1
		for left <= right {
			mid := left + (right-left)>>1
			if t-this.Cache[mid] > 3000 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		this.Cache = append(this.Cache[left:], t)
	}
	return len(this.Cache)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
