package main

func main() {

}

func numJewelsInStones(jewels string, stones string) int {
	cache := make(map[rune]bool)
	for _, j := range jewels {
		cache[j] = true
	}
	ans := 0
	for _, s := range stones {
		if cache[s] {
			ans++
		}
	}
	return ans
}
