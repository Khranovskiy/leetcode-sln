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
