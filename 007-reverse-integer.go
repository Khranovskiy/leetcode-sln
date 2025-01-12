package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse2(123))
	fmt.Println(reverse2(120))
	fmt.Println(reverse2(-123))
	fmt.Println(reverse2(math.MaxInt32))
	//               2147483647    -> 7463847412 -> 0
	//      ?     -> 2144847412    -> 2147484412
	//          (max 2147483647) (max 2147483647)
	fmt.Println(reverse2(2144847412))
}

func reverse2(left int) int {
	res := 0

	for left != 0 {
		right := left % 10
		left = left / 10

		if res > math.MaxInt32/10 || res == math.MaxInt32/10 && right > 7 {
			return 0
		}
		if res < math.MinInt32/10 || res == math.MinInt32/10 && right < -7 {
			return 0
		}
		res = res*10 + right
	}
	return res
}
