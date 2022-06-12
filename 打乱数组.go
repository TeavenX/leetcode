package main

import "math/rand"

func main() {

}

type Solution struct {
	bucket []int
	length int
}

func Constructor(nums []int) Solution {
	return Solution{
		bucket: nums,
		length: len(nums),
	}
}

func (this *Solution) Reset() []int {
	return this.bucket
}

func (this *Solution) Shuffle() []int {
	result := make([]int, this.length)
	copy(result, this.bucket)
	for i := 0; i < this.length; i++ {
		j := rand.Intn(this.length)
		result[i], result[j] = result[j], result[i]
	}
	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
