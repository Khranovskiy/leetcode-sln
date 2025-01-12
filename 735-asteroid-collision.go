package main

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for i := 0; i < len(asteroids); i++ {
		stack = appendToStack(stack, asteroids[i])
	}
	return stack
}
func appendToStack(stack []int, current int) []int {
	if len(stack) == 0 {
		return append(stack, current)
	}
	top := stack[len(stack)-1]

	if top*current > 0 || top < 0 {
		// put back to the stack
		// add the current element to the stack.
		return append(stack, current)
	}
	// top positive and curr element negative
	// curr asteroid has a smaller mass
	if top > -1*current {
		// return top back to the stack and move to the next element.
		return stack
	}

	// If the top asteroid has a smaller mass, we extract the next element from the stack without returning the top element to the stack and repeat the checks with a new pair.
	if top < -1*current {
		return appendToStack(stack[:len(stack)-1], current)
	}
	// If the elements have the same mass
	// move to the next element
	// without returning the top element to the stack.
	return stack[:len(stack)-1]
}
