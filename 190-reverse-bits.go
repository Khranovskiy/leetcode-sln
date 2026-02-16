package main

import "fmt"

// https://leetcode.com/problems/reverse-bits/?envType=daily-question&envId=2026-02-16

func main() {
	fmt.Println(reverseBits(1))
}

// For each of the 32 iterations:

// result << 1 — shift the result left to make room for the next bit.
// num & 1 — extract the lowest bit of num.
// | — place that bit into the lowest position of result.
// num >>= 1 — shift num right to process the next bit.

func reverseBitsSlow(num int) int {
	var result int
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}

var cache [256]int

func initBetter() {
	for i := 0; i < 256; i++ {
		var rev int
		val := i
		for j := 0; j < 8; j++ {
			rev = (rev << 1) | (val & 1)
			val >>= 1
		}
		cache[i] = rev
	}
}

func reverseBitsBetter(num int) int {
	return cache[num&0xff]<<24 |
		cache[(num>>8)&0xff]<<16 |
		cache[(num>>16)&0xff]<<8 |
		cache[(num>>24)&0xff]
}

// best
func reverseBits(num int) int {
	num = (num>>16)&0x0000FFFF | (num&0x0000FFFF)<<16
	num = (num>>8)&0x00FF00FF | (num&0x00FF00FF)<<8
	num = (num>>4)&0x0F0F0F0F | (num&0x0F0F0F0F)<<4
	num = (num>>2)&0x33333333 | (num&0x33333333)<<2
	num = (num>>1)&0x55555555 | (num&0x55555555)<<1
	return num
}
