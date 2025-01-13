package main

import "fmt"

func main() {
	fmt.Println(isSameAfterReversals(526))
}

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	return !(num%10 == 0)
}

func reverse(num int) int {
	res := 0

	for num != 0 {
		right := num % 10
		num = num / 10

		res = res*10 + right
	}

	return res
}
