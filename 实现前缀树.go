package main

func main() {

}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	t := this
	for _, b := range word {
		if t.next[b-'a'] == nil {
			t.next[b-'a'] = &Trie{}
		}
		t = t.next[b-'a']
	}
	t.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for _, b := range word {
		if t.next[b-'a'] == nil {
			return false
		}
		t = t.next[b-'a']
	}
	return t.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for _, b := range prefix {
		if t.next[b-'a'] == nil {
			return false
		}
		t = t.next[b-'a']
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
