package main

import "fmt"

func main() {
	// nums := []int{1, 4, 2, 7}
	// fmt.Println(countPairs(nums, 2, 6))
	nums := []int{9, 8, 4, 2, 1}
	fmt.Println(countPairs(nums, 5, 14))
}

func countPairsTME(nums []int, low int, high int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if r := nums[i] ^ nums[j]; r >= low && r <= high {
				count++
			}
		}
	}
	return count
}

type Trie struct {
	ch  [2]*Trie
	siz int
}

func (t *Trie) insert(num int) {
	tmp := t
	for i := 15; i >= 0; i-- {
		c := num >> i & 1
		if tmp.ch[c] == nil {
			tmp.ch[c] = &Trie{}
		}
		tmp.siz++
		tmp = tmp.ch[c]
	}
	tmp.siz++
}

func (t *Trie) query(num int, max int) int {
	tmp := t
	result := 0
	for i := 15; i >= 0 && tmp != nil; i-- {
		cn := num >> i & 1
		cm := max >> i & 1
		next := -1
		for c := 0; c < 2; c++ {
			if tmp.ch[c] == nil {
				continue
			}
			if c^cn < cm {
				result += tmp.ch[c].siz
			}
			if c^cn == cm {
				next = c
			}
		}
		if next != -1 && tmp.ch[next] != nil {
			tmp = tmp.ch[next]
		} else {
			break
		}
	}
	return result
}

func countPairs(nums []int, low int, high int) int {
	trie := &Trie{}
	ans := 0
	for _, num := range nums {
		ans += trie.query(num, high+1) - trie.query(num, low)
		trie.insert(num)
	}
	return ans
}
