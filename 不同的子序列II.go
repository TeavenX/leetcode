package main

func main() {
	distinctSubseqII("abb")
}

const mod = 1e9 + 7

func distinctSubseqII(s string) int {
	/*
	   > aba
	   {} <------------------------+       哨兵
	   {}, a                       |       a           pre = [{}]              len(pre) = 1
	   {}, a, ab, b                |       b           pre = [{}, a]           len(pre) = 2
	   {}, a, ab, b, aa, aba, ba, *a       a           pre = [{}, a, ab, b]    len(pre) = 4

	   > abaa
	   {} <------------------------+                                    哨兵
	   {}, a                       |                                    a           pre = [{}]              len(pre) = 1
	   {}, a, ab, b                |                                    b           pre = [{}, a]           len(pre) = 2
	   {}, a, ab, b, aa, aba, ba, *a                                    a           pre = [{}, a, ab, b]    len(pre) = 4
	   {}, a, ab, b, aa, aba, ba, *aa, *aba, *ba, aaa, abaa, baa, a     a           pre = [{}, a, ab, b, aa, aba, ba]
	*/
	// 保存相同字符结尾的重复值，因为cur = pre（之前的值） + pre（之前值组合当前的字符）- 重复值
	cache := [26]int{}
	// 设置一个哨兵
	pre, cur := 1, 0
	for _, c := range s {
		i := c - 'a'
		// 从缓存中取之前重复过的（相同字符结尾的值）
		dup := cache[i]
		// 执行上面的公式
		cur = (pre + pre - dup + mod) % mod
		// 这里存pre的原因是，当前遍历的字符串已经把「前面的字符串(pre)」都组合过一遍了
		// 下一次遇到时，需要把之前组合过的部分给删掉
		cache[i] = pre
		pre = cur
	}
	// 把之前设置的哨兵取消
	return cur - 1
}
