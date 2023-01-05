package main

func main() {

}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	t := this
	for _, b := range word {
		i := b - 'a'
		if t.children[i] == nil {
			t.children[i] = &Trie{}
		}
		t = t.children[i]
	}
	t.isEnd = true
}

func (this *Trie) Search(word string) bool {
	t := this
	for _, b := range word {
		i := b - 'a'
		if t.children[i] == nil {
			return false
		}
		t = t.children[i]
	}
	return t.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for _, b := range prefix {
		i := b - 'a'
		if t.children[i] == nil {
			return false
		}
		t = t.children[i]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
