package main

func main() {

}

func equalFrequency(word string) bool {
	cache := make(map[rune]int)
	for _, b := range word {
		cache[b]++
	}
	cnt := make([]int, 0, len(cache))
	for _, c := range cache {
		cnt = append(cnt, c)
	}
	sort.Ints(cnt)
	n := len(cnt)
	return n == 1 ||
		cnt[0] == 1 && same(cnt[1:]) ||
		cnt[n-1] == cnt[n-2]+1 && same(cnt[:n-1])
}

func same(a []int) bool {
	for _, v := range a[1:] {
		if v != a[0] {
			return false
		}
	}
	return true
}
