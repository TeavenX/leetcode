package main

func main() {

}

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	left, right := 0, n-1
	for 0 < right && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 {
		return 0
	}
	ans := right
	for ; left == 0 || arr[left-1] <= arr[left]; left++ {
		for right < n && arr[left] > arr[right] {
			right++
		}
		ans = min(ans, right-left-1)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
