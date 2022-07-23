package main

func main() {

}

type degree struct {
	x, y int
}

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	set := make(map[degree]bool)
	for _, seq := range sequences {
		for i := 0; i < len(seq)-1; i++ {
			set[degree{seq[i], seq[i+1]}] = true
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if exist := set[degree{nums[i], nums[i+1]}]; !exist {
			return false
		}
	}
	return true
}
