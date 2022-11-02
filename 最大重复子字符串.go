package main

func main() {

}

func maxRepeating(sequence string, word string) int {
	n := len(word)
	m := len(sequence)
	max := 0
	for i := 0; i <= m-n; i++ {
		t := i
		count := 0
		for t+n <= m && word == sequence[t:t+n] {
			count++
			t += n
			if max < count {
				max = count
			}
		}
	}
	return max
}
