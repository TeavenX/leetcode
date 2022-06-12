package main

func main() {

}

func minStickers(stickers []string, target string) int {
	n := len(stickers)
	dict := make([]map[rune]int, n)
    source := make(map[rune]int, len(target))
	for idx, sticker := range stickers {
		for _, str := range sticker {
			dict[idx][str]++
		}
	}
    for _, str := range target {
        source[str]++
    }
    for k, v := range source {
        for temp := range source {
            if temp[k] > v
        }
    }
}
