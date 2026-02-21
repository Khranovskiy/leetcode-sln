package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(97)
	result := n.ProbablyPrime(20) // 20 = number of rounds
	fmt.Println(result)           // true

	fmt.Println(countPrimeSetBits(10, 15))
}

func countPrimeSetBits(left int, right int) int {
	result := 0

	for i := left; i <= right; i++ {
		bits := countBits(i)
		isPrime := big.NewInt(int64(bits)).ProbablyPrime(10)
		if isPrime {
			result++
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
