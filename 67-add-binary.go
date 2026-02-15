package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\"0\"+ \"0\" =  \"%s\"", addBinary("0", "0"))
}

func addBinary(a string, b string) string {
	result := make([]byte, max(len(a), len(b))+1)

	i, j := len(a)-1, len(b)-1
	k := len(result) - 1
	carry := byte(0)

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}

		result[k] = (sum % 2) + '0'
		carry = sum / 2
		k--
	}
	for k >= 0 {
		result[k] = '0'
		k--
	}
	start := 0
	for start < len(result)-1 && result[start] == '0' {
		start++
	}

	return string(result[start:])
}
