package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(3))
}
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			res[i-1] = "Fizz"
		}
		if i%5 == 0 {
			res[i-1] = res[i-1] + "Buzz"
		}
		if len(res[i-1]) == 0 {
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res
}
