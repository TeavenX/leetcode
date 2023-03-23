package main

func main() {

}

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

func (t *Trie) insert(val string) {
	n := len(val)
	node := t
	for i := n - 1; i >= 0; i-- {
		b := val[i] - 'a'
		if node.next[b] == nil {
			node.next[b] = &Trie{}
		}
		node = node.next[b]
	}
	node.isEnd = true
}

func (t *Trie) search(stream []byte) bool {
	node := t
	for i := len(stream) - 1; i >= 0; i-- {
		b := stream[i] - 'a'
		if node.next[b] == nil {
			return false
		}
		if node.next[b].isEnd {
			return true
		}
		node = node.next[b]
	}
	return false
}

type StreamChecker struct {
	root   *Trie
	stream []byte
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{root: &Trie{}, stream: make([]byte, 0)}
	for _, word := range words {
		sc.root.insert(word)
	}
	return sc
}

func (this *StreamChecker) Query(letter byte) bool {
	this.stream = append(this.stream, letter)
	return this.root.search(this.stream)
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
