package main

func main() {

}

func numberOfCuts(n int) int {
	if n > 1 && n&1 == 1 {
		return n
	}
	return n >> 1
}
