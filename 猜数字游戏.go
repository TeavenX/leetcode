package main

import "fmt"

func main() {

}

func getHint(secret string, guess string) string {
	scache := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		scache[secret[i]]++
	}
	rightCount, wrongCount := 0, 0
	mismatch := make([]int, 0)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			scache[secret[i]]--
			rightCount++
		} else {
			mismatch = append(mismatch, i)
		}
	}
	for _, i := range mismatch {
		if count, ok := scache[guess[i]]; ok {
			if count > 0 {
				scache[guess[i]]--
				wrongCount++
			}
		}
	}
	return fmt.Sprintf("%dA%dB", rightCount, wrongCount)
}

func getHintV2(secret string, guess string) string {
	scache := [10]int{}
	for i := 0; i < len(secret); i++ {
		scache[secret[i]-'0']++
	}
	rightCount, wrongCount := 0, 0
	mismatch := make([]int, 0)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			scache[secret[i]-'0']--
			rightCount++
		} else {
			mismatch = append(mismatch, i)
		}
	}
	for _, i := range mismatch {
		if count := scache[guess[i]-'0']; count > 0 {
			scache[guess[i]-'0']--
			wrongCount++
		}
	}
	return fmt.Sprintf("%dA%dB", rightCount, wrongCount)
}
