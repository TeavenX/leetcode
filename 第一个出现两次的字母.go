package main

func main() {

}

func repeatedCharacter(s string) byte {
	cache := 0
	for _, b := range s {
		i := b - 'a'
		if (cache>>i)&1 == 1 {
			return byte(b)
		}
		cache |= 1 << i
	}
	return 'a'
}
