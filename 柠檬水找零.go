package main

func main() {

}

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		if b == 20 {
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		} else if b == 10 {
			ten++
			five--
		} else {
			five++
		}
		if five < 0 {
			return false
		}
	}
	return true
}
