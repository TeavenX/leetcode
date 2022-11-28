package main

func main() {

}

func minOperations(s string) int {
	zero := 0
	for i, b := range s {
		if i&1 == 0 {
			if b == '1' {
				zero++
			}
		} else {
			if b == '0' {
				zero++
			}
		}
	}
	return min(zero, len(s)-zero)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
