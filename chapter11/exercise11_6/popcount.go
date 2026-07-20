package exercise116

// numbers from 0->255 (1 byte) has how many 1s each
// pc[23] is number of 1 bits of 23 (integer)
var pc [256]byte

func init() {
	// i = 0,1,..,255
	// "i/2" is "i shifting right 1"
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// split a long 8-byte number into 8 parts; count number of 1s of each part
func PopCount1(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var res int
	for i := 0; i < 64; i++ {
		res += int(byte(x>>i) & 1)
	}
	return res
}

// using the fact that x&(x-1) turns off the last "1" bit to "0"
func PopCount3(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
