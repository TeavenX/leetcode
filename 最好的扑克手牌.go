package main

func main() {

}

func bestHand(ranks []int, suits []byte) string {
	cn := [14]int{}
	cs := ['e']byte{}
	for i, r := range ranks {
		cn[r]++
		cs[suits[i]]++
		switch {
		case cs[suits[i]] == 5:
			return "Flush"
		case cn[r] == 3:
			return "Three of a Kind"
		}
	}
	for _, c := range cn {
		if c == 2 {
			return "Pair"
		}
	}
	return "High Card"
}
