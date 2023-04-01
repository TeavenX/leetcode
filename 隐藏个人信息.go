package main

import (
	"bytes"
	"strings"
)

func main() {

}

func maskPII(s string) string {
	if strings.Contains(s, "@") {
		return maskEmail(s)
	}
	return maskPhone(s)
}

func maskEmail(s string) string {
	buf := bytes.Buffer{}
	b := s[0]
	if b <= 'Z' {
		b = b - 'A' + 'a'
	}
	buf.WriteByte(b)
	buf.Write([]byte{'*', '*', '*', '*', '*'})
	flag := false
	for i := 1; i < len(s); i++ {
		b := s[i]
		if b >= 'A' && b <= 'Z' {
			b = b - 'A' + 'a'
		}
		if b == '@' {
			c := s[i-1]
			if c <= 'Z' && c >= 'A' {
				c = c - 'A' + 'a'
			}
			buf.WriteByte(c)
			flag = true
		}
		if flag {
			buf.WriteByte(b)
		}
	}
	return buf.String()
}

func maskPhone(s string) string {
	c := make([]byte, 32)
	count := 0
	idx := 31
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			if idx > len(c)-5 {
				c[idx] = s[i]
			} else {
				c[idx] = '*'
			}
			idx--
			count++
			if count > 3 && (count-4)%3 == 0 {
				c[idx] = '-'
				idx--
			}
		}
	}
	if count > 10 {
		if (count-4)%3 == 0 {
			idx++
		}
		c[idx] = '+'
		return string(c[idx:])
	}
	return string(c[idx+2:])
}
