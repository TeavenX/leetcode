package main

func main() {

}

func canFormArray(arr []int, pieces [][]int) bool {
	for i := 0; i < len(arr); i++ {
		result := false
		for _, piece := range pieces {
			if piece[0] == arr[i] {
				result = true
				for j := 1; j < len(piece); j++ {
					i++
					if piece[j] != arr[i] {
						return false
					}
				}
			}
		}
		if !result {
			return false
		}
	}
	return true
}
