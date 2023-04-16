package main

func main() {

}

type MajorityChecker []int

func Constructor(arr []int) MajorityChecker {
	return MajorityChecker(arr)
}

func (this *MajorityChecker) QueryTLE(left int, right int, threshold int) int {
	cache := [20001]int{}
	for i := left; i <= right; i++ {
		cache[(*this)[i]]++
		if cache[(*this)[i]] >= threshold {
			return (*this)[i]
		}
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
