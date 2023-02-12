package main

import "bytes"

func main() {

}

func alphabetBoardPath(target string) string {
	buf := bytes.Buffer{}
	i, j := 0, 0
	for _, b := range target {
		t := int(b - 'a')
		x, y := t/5, t%5
		if x == i && y == j {
			buf.WriteByte('!')
			continue
		}
		dx, dy := x-i, y-j
		for ; dy < 0; dy++ {
			buf.WriteByte('L')
			j--
		}
		for ; dx > 0; dx-- {
			buf.WriteByte('D')
			i++
		}
		for ; dx < 0; dx++ {
			buf.WriteByte('U')
			i--
		}
		for ; dy > 0; dy-- {
			buf.WriteByte('R')
			j++
		}
		buf.WriteByte('!')
	}
	return buf.String()
}
