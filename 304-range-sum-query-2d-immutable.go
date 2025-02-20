package main

import "fmt"

type NumMatrix struct {
	prefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	if rows == 0 {
		return NumMatrix{}
	}
	cols := len(matrix[0])

	prefixSum := make([][]int, rows+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, cols+1)
	}

	// Build the prefix sum array
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			prefixSum[i][j] = matrix[i-1][j-1] +
				prefixSum[i-1][j] + // Top
				prefixSum[i][j-1] - // Left
				prefixSum[i-1][j-1] // Top-left
		}
	}

	return NumMatrix{prefixSum: prefixSum}
}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	// Shift indices to account for the prefixSum offset
	row1, col1, row2, col2 = row1+1, col1+1, row2+1, col2+1

	return this.prefixSum[row2][col2] -
		this.prefixSum[row1-1][col2] -
		this.prefixSum[row2][col1-1] +
		this.prefixSum[row1-1][col1-1]
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	obj := Constructor(matrix)
	fmt.Println(obj.SumRegion(2, 1, 4, 3)) // Output: 8
	fmt.Println(obj.SumRegion(1, 1, 2, 2)) // Output: 11
	fmt.Println(obj.SumRegion(1, 2, 2, 4)) // Output: 12
}
