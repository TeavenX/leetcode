package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func getFolderNamesError(names []string) []string {
	cache := make(map[string]int)
	ans := make([]string, len(names))
	for idx, name := range names {
		if v, ok := cache[name]; ok {
			i := 0
			for v>>i&1 == 1 {
				i++
			}
			cache[name] |= 1 << i
			ans[idx] = fmt.Sprintf("%s(%d)", name, i)
			cache[ans[idx]] = 1
		} else {
			num := 0
			j := 0
			for i := len(name) - 1; i >= 0; i-- {
				b := name[i]
				if b == '(' {
					name = name[:i]
					break
				} else if '0' <= b && b <= '9' {
					num += j*10 + int(b-'0')
				}
			}
			j = 0
			for cache[name]>>j&1 == 1 {
				j++
			}
			cache[name] |= 1 << j
			if j == 0 {
				ans[idx] = name
			} else {
				ans[idx] = fmt.Sprintf("%s(%d)", name, j)
			}
		}
	}
	return ans
}

func getFolderNames(names []string) []string {
	cache := make(map[string]int)
	ans := make([]string, len(names))
	for idx, name := range names {
		i := cache[name]
		if i == 0 {
			ans[idx] = name
			cache[name] = 1
			continue
		}
		// 用fmt.Sprintf慢了40ms，内存开销增加2mb
		for cache[name+"("+strconv.Itoa(i)+")"] > 0 {
			i++
		}
		cache[name] = i + 1
		name = name + "(" + strconv.Itoa(i) + ")"
		cache[name] = 1
		ans[idx] = name
	}
	return ans
}
