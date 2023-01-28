package main

import "fmt"

func main() {
	nums := []int{2, 1, 6, 4}
	fmt.Println(waysToMakeFair(nums))
}

func waysToMakeFairError(nums []int) int {
	n := len(nums)
	odd := make([]int, 0, n>>1+1)
	even := make([]int, 0, n>>1+1)
	for i, num := range nums {
		if i&1 == 0 {
			if i > 1 {
				num += even[len(even)-1]
			}
			even = append(even, num)
		} else {
			if i > 1 {
				num += odd[len(odd)-1]
			}
			odd = append(odd, num)
		}
	}
	ans := 0
	for i, num := range nums {
		idx := i >> 1
		e := even[idx]
		o := odd[idx]
		if i&1 == 0 {
			e -= num
		} else {
			o -= num
		}
		e += odd[len(odd)-1] - odd[idx]
		o += even[len(even)-1] - even[idx]
		if o == e {
			ans++
		}
	}
	return ans
}

func waysToMakeFair(nums []int) int {
	var s1, s2, t1, t2 int
	for i, num := range nums {
		if i&1 == 0 {
			s1 += num
		} else {
			s2 += num
		}
	}
	ans := 0
	for i, num := range nums {
		if i&1 == 0 && t1+s2-t2 == t2+s1-t1-num {
			ans++
		}
		if i&1 == 1 && t1+s2-t2-num == t2+s1-t1 {
			ans++
		}
		if i&1 == 0 {
			t1 += num
		} else {
			t2 += num
		}
	}
	return ans
}
