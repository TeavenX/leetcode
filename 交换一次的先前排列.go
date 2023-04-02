package main

func main() {

}

func prevPermOpt1(arr []int) []int {
	n := len(arr)
	var maxVal, maxIdx int
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			maxVal = arr[i+1]
			maxIdx = i + 1
			for j := i + 2; j < n; j++ {
				if arr[j] > maxVal && arr[j] < arr[i] {
					maxVal = arr[j]
					maxIdx = j
				}
			}
			arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
			break
		}
	}
	return arr
}
