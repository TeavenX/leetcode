package main

func main() {

}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	nF, nS := len(firstList), len(secondList)
	result := make([][]int, 0)
	if nF == 0 || nS == 0 {
		return result
	}
	ptrF, ptrS := 0, 0
	for ptrF < nF && ptrS < nS {
		start := max(firstList[ptrF][0], secondList[ptrS][0])
		end := min(firstList[ptrF][1], secondList[ptrS][1])
		if start <= end {
			result = append(result, []int{start, end})
		}
		if firstList[ptrF][1] < secondList[ptrS][1] {
			ptrF++
		} else {
			ptrS++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
