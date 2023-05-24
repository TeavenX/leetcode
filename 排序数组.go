package main

import "math/rand"

func main() {

}

func sortArray(nums []int) []int {
	var sort func(left, right int)
	sort = func(left, right int) {
		if left >= right {
			return
		}
		pos := left + rand.Intn(right-left+1)
		nums[left], nums[pos] = nums[pos], nums[left]
		j := left
		for i := left + 1; i <= right; i++ {
			if nums[i] > nums[left] {
				continue
			}
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[left], nums[j] = nums[j], nums[left]
		sort(left, j-1)
		sort(j+1, right)
	}
	sort(0, len(nums)-1)
	return nums
}

func sortArrayV2(nums []int) []int {
	var sort func(left, right int)
	sort = func(left, right int) {
		if left >= right {
			return
		}
		pos := left + rand.Intn(right-left+1)
		nums[left], nums[pos] = nums[pos], nums[left]

		less, greater := left+1, right
		for {
			for less <= greater && nums[less] < nums[left] {
				less++
			}
			for less <= greater && nums[greater] > nums[left] {
				greater--
			}
			if less >= greater {
				break
			}
			nums[less], nums[greater] = nums[greater], nums[less]
			less++
			greater--
		}

		nums[left], nums[greater] = nums[greater], nums[left]
		sort(left, greater-1)
		sort(greater+1, right)
	}
	sort(0, len(nums)-1)
	return nums
}

func sortArrayV3(nums []int) []int {
	var sort func(left, right int)
	sort = func(left, right int) {
		if left >= right {
			return
		}
		pos := left + rand.Intn(right-left+1)
		nums[left], nums[pos] = nums[pos], nums[left]

		pivot := nums[left]

		le, ge := left+1, right
		for i := left + 1; i <= ge; {
			if nums[i] > pivot {
				nums[ge], nums[i] = nums[i], nums[ge]
				ge--
			} else if nums[i] < pivot {
				nums[le], nums[i] = nums[i], nums[le]
				le++
				i++
			} else {
				i++
			}
		}

		nums[left], nums[le-1] = nums[le-1], nums[left]
		sort(left, le-2)
		sort(ge+1, right)
	}
	sort(0, len(nums)-1)
	return nums
}
