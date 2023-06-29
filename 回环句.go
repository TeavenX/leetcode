package main

func main() {

}

func isCircularSentence(sentence string) bool {
	for i, b := range sentence {
		if b == ' ' {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
		}
	}
	return sentence[0] == sentence[len(sentence)-1]
}
