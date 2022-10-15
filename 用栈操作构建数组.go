package main

func main() {

}

func buildArray1(target []int, n int) []string {
	result := make([]string, 0, 2*n-len(target)) // 4ms
	idx := 1
	for _, num := range target {
		for idx < num {
			result = append(result, "Push", "Pop")
			idx++
		}
		result = append(result, "Push")
		idx++
	}
	return result
}

func buildArray2(target []int, n int) []string {
	cap := 2*n - len(target)
	result := make([]string, 0, cap) // 0ms
	idx := 1
	for _, num := range target {
		for idx < num {
			result = append(result, "Push", "Pop")
			idx++
		}
		result = append(result, "Push")
		idx++
	}
	return result
}
