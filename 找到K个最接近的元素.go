package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{1, 2, 3, 4, 5}
	//arr := []int{2, 4, 6, 8, 10}
	//arr := []int{1, 1, 1, 10, 10, 10}
	//fmt.Println(findClosestElements(arr, 1, 9))
	//arr := []int{1, 2, 3, 4, 4, 4, 4, 5, 5}
	//fmt.Println(findClosestElements(arr, 3, 3))
	//arr := []int{1, 10, 15, 25, 35, 45, 50, 59}
	arr := []int{0, 1, 2, 2, 2, 3, 6, 8, 8, 9}
	fmt.Println(findClosestElementsV1(arr, 5, 9))
}

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := (left + right) >> 1
		if arr[mid] == x {
			left = mid
			right = mid
			break
		} else if arr[mid] < x && abs(arr[right]-x) < abs(arr[left]-x) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	k -= right - left + 1
	result := append([]int{}, arr[left:right+1]...)
	left--
	right++
	for k > 0 {
		if left < 0 {
			result = append(result, arr[right])
			right++
		} else if right >= len(arr) {
			result = append(result, arr[left])
			left--
		} else {
			l, r := abs(arr[left]-x), abs(arr[right]-x)
			if l <= r {
				result = append(result, arr[left])
				left--
			} else {
				result = append(result, arr[right])
				right++
			}
		}
		k--
	}
	sort.Ints(result)
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findClosestElementsV1(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	}
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	left, right := 0, len(arr)-k
	for left < right {
		mid := (left + right) >> 1
		if arr[mid] == x {
			left = mid
			break
		} else if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
