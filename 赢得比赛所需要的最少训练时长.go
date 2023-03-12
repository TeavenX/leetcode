package main

func main() {

}

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	ans := 0
	totalEng := 0
	for i, exp := range experience {
		totalEng += energy[i]
		if initialExperience <= exp {
			ans += exp - initialExperience + 1
			initialExperience = exp + 1
		}
		initialExperience += exp
	}
	if totalEng >= initialEnergy {
		ans += totalEng - initialEnergy + 1
	}
	return ans
}
