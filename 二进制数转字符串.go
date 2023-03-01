package main

import "bytes"

func main() {

}

func printBin(num float64) string {
	b := bytes.Buffer{}
	b.WriteString("0.")
	for b.Len() < 32 {
		num = num * 2
		if num >= 1 {
			b.WriteByte('1')
			num--
			if num == 0 {
				return b.String()
			}
		} else {
			b.WriteByte('0')
		}
	}
	return "ERROR"
}
