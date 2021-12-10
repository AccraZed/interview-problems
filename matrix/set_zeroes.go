package matrix

import "fmt"

// https://leetcode.com/problems/set-matrix-zeroes/

func SetZeroesMain() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	fmt.Printf("Before: %v\n", matrix)
	setZeroes(matrix)
	fmt.Printf("After: %v\n", matrix)
}

// Space Complexity: O(m + n) 
// Time Complexity: O(2*m*n) -> O(m*n)
func setZeroes(matrix [][]int) {
	mRows := make([]bool, len(matrix))
	nRows := make([]bool, len(matrix[0]))

	// Find zeroes
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				mRows[i] = true
				nRows[j] = true
			}
		}
	}

	// Set row zeroes
	for i := 0; i < len(matrix); i++ {
		if mRows[i] {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	// Set col zeroes
	for j := 0; j < len(matrix[0]); j++ {
		if nRows[j] {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
}
