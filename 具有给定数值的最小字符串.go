package main

func main() {

}

func getSmallestString(n int, k int) string {
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = 'a'
	}
	k -= n
	for k > 0 {
		idx := n - 1
		for ; k > 25; k -= 25 {
			ans[idx] = 'z'
			idx--
		}
		ans[idx] = ans[idx] + byte(k)
		k = 0
	}
	return string(ans)
}
