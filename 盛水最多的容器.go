package main

func main() {

}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		var temp int
		if height[left] < height[right] {
			temp = height[left] * (right - left)
			left++
		} else {
			temp = height[right] * (right - left)
			right--
		}
		if temp > result {
			result = temp
		}
	}
	return result
}
