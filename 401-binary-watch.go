package main

import "fmt"

func main() {
	fmt.Println(readBinaryWatch(1))
}

func readBinaryWatch(turnedOn int) []string {
	// 0-11.             1 0 1 1 -> 0..3 leds
	// 0-59.    0 0 1 1  1 0 1 1 -> 0..5 leds

	var result = make([]string, 0)

	for h := 0; h <= 11; h++ {
		for m := 0; m <= 59; m++ {
			if countBits(h)+countBits(m) == turnedOn {
				result = append(result, fmt.Sprintf("%01d:%02d", h, m))
			}
		}
	}
	return result
}

func countBits(n int) int {
	var res int
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
