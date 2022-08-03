package main

func main() {

}

func generateTheString(n int) string {
	result := make([]byte, n)
	if n&1 == 0 {
		for i := 0; i < n-1; i++ {
			result[i] = 'a'
		}
		result[n-1] = 'b'
	} else {
		for i := 0; i < n; i++ {
			result[i] = 'a'
		}
	}
	return string(result)
}
