package main

import "fmt"

func main() {
	m := Constructor()
	m.BuildDict([]string{"a"})
	fmt.Println(m.Search("b"))
}

type MagicDictionary struct {
	m map[int][]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		m: make(map[int][]string),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.m[len(word)] = append(this.m[len(word)], word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
next:
	for _, word := range this.m[len(searchWord)] {
		diff := 0
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				diff++
				if diff > 1 {
					continue next
				}
			}
		}
		if diff == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
