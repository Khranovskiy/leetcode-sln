package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures)) // filled with zeros
	for j, dailyTemp := range temperatures {
		for len(stack) > 0 && dailyTemp > temperatures[stack[len(stack)-1]] {
			i := stack[len(stack)-1] // индекс предыдущего дня попрохладне
			stack = stack[:len(stack)-1]
			res[i] = j - i // j or tempIndex индекс первого теплого дня
		}
		stack = append(stack, j)
	}
	return res
}

// [73,74,75,71,69,72,76,73]

// [9 8 7 6 5 4 3 2 10]
// [8 7 6 5 4 3 2 1 0]

// [ 1 2 3 4 5 6 7 8]

// j = 0, dayTemp 1, stack []: stack is [0]
// j = 1, dayTemp 2, stack [0]: 2 > temp[ stack head aka 0] 2 > 1, stack [], res[head] = j - head
