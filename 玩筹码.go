package main

func main() {

}

func minCostToMoveChips(position []int) int {
	numEven := 0
	numOdd := 0
	for _, num := range position {
		if num&1 == 0 {
			numEven++
		} else {
			numOdd++
		}
	}
	if numEven < numOdd {
		return numEven
	}
	return numOdd
}
