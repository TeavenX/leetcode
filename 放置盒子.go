package main

func main() {

}

// 20ms 1.9MB
func minimumBoxes(n int) int {
	c := 1
	b := 0
	box := 0
	for box < n {
		for i := 1; i <= c && box < n; i++ {
			box += i
			b++
		}
		c++
	}
	return b
}

// 0ms 1.9MB
func minimumBoxesV2(n int) int {
	c := 1
	b := 0
	box := 0
	for box+b+c < n {
		b += c
		box += b
		c++
	}
	for i := 1; i <= c && box < n; i++ {
		box += i
		b++
	}
	return b
}
