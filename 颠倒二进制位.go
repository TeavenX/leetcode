package main

func main() {

}

const (
	m1 = 0x55555555 // 01010101010101010101010101010101
	m2 = 0x33333333 // 00110011001100110011001100110011
	m4 = 0x0f0f0f0f // 00001111000011110000111100001111
	m8 = 0x00ff00ff // 00000000111111110000000011111111
)

func reverseBits(num uint32) uint32 {
	num = num>>1&m1 | num&m1<<1
	num = num>>2&m2 | num&m2<<2
	num = num>>4&m4 | num&m4<<4
	num = num>>8&m8 | num&m8<<8
	return num<<16 | num>>16
}

func reverseBitsV1(num uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 32 && num > 0; i++ {
		result |= num & 1 << (31 - i)
		num >>= 1
	}
	return result
}
