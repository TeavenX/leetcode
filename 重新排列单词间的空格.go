package main

func main() {

}

func reorderSpaces(text string) string {
	if len(text) == 1 {
		return text
	}
	scount, wcount := 0, 0
	if text[0] != ' ' {
		wcount++
	}
	for i, s := range text {
		if s == ' ' {
			scount++
		} else if i > 0 && text[i-1] == ' ' {
			wcount++
		}
	}
	perSpace := 0
	if wcount-1 > 0 {
		perSpace = scount / (wcount - 1)
	}
	result := make([]byte, 0, len(text))
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			continue
		}
		for i < len(text) && text[i] != ' ' {
			result = append(result, text[i])
			i++
		}
		for j := 0; j < perSpace && len(result) > 0; j++ {
			result = append(result, ' ')
		}
	}
	for len(result) < len(text) {
		result = append(result, ' ')
	}
	return string(result[:len(text)])
}
