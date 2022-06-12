package main

import "fmt"

func main() {
	a := []int{0, 1, 0, 3, 12}
	a = []int{1, 0}
	a = []int{1}
	a = []int{0, 0, 0, 0}
	a = []int{1, 1, 1, 1}
	a = []int{1, 0, 0, 1}
	moveZeroesLeetCode(a)
	fmt.Println(a)
}

func moveZeroesBeauty(nums []int) {
	slow, fast, lenN := 0, 0, len(nums)
	for fast < lenN {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
}

func moveZeroesLeetCode(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func moveZeroes(nums []int) {
	slow, fast, lenN := 0, 0, len(nums)
	for fast < lenN {
		if nums[fast] == 0 {
			fast += 1
		} else {
			if nums[slow] != 0 {
				slow += 1
				if slow > fast {
					fast = slow
				}
			} else {
				nums[fast], nums[slow] = nums[slow], nums[fast]
			}
		}
	}
}
