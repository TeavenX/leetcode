package main

func main() {

}

func defangIPaddr(address string) string {
	arr := []byte(address)
	result := make([]byte, 0, len(arr)+6)
	for _, str := range arr {
		if str == '.' {
			result = append(result, '[')
			result = append(result, '.')
			result = append(result, ']')
		} else {
			result = append(result, str)
		}
	}
	return string(result)
}

func defangIPaddrV2(address string) string {
	arr := make([]byte, 0, len(address)+6)
	for _, str := range address {
		if str == '.' {
			arr = append(arr, '[', '.', ']')
		} else {
			arr = append(arr, uint8(str))
		}
	}
	return string(arr)
}
