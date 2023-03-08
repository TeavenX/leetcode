package main

func main() {

}

func minimumRecolors(blocks string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			count++
		}
	}
	ans := count
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			count++
		}
		if blocks[i-k] == 'W' {
			count--
		}
		if count < ans {
			ans = count
		}
	}
	return ans
}
