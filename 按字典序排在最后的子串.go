package main

import (
	"fmt"
	"index/suffixarray"
	"reflect"
	"unsafe"
)

func main() {
	lastSubstringV2("abab")
}

func lastSubstring(s string) string {
	mx := 'a'
	ans := ""
	for i, b := range s {
		if b > mx {
			mx = b
		}
		if b == mx {
			if v := s[i:]; v > ans {
				ans = v
			}
		}
	}
	return ans
}

func lastSubstringV2(s string) string {
	n := len(s)

	sfa := suffixarray.New([]byte(s))
	fmt.Println(sfa)

	sfaAddr := reflect.ValueOf(sfa).Elem().FieldByName("sa").Field(0)

	sa := *(*[]int32)(unsafe.Pointer(sfaAddr.UnsafeAddr()))
	fmt.Println(sa)
	return s[sa[n-1]:]
}
