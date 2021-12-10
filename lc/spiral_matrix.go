package lc

import "fmt"

func SpiralMatrixMain() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 5},
		{9, 8, 7, 6},
	}

	matrix2 := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}

	fmt.Println(spiralOrder(matrix))
	fmt.Println(spiralOrder(matrix2))
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	rows, cols := m, n
	i, j, k := 0, 0, 0

	res := make([]int, 0, m*n)
	for rows > 0 && cols > 0 {
		// up -> right
		for k = 0; k < cols; k++ {
			res = append(res, matrix[i][j])
			j++
		}
		j--
		i++
		rows--
		// right -> down
		for k = 0; k < rows; k++ {
			res = append(res, matrix[i][j])
			i++
		}
		i--
		j--
		cols--
		// down -> left
		for k = 0; k < cols; k++ {
			res = append(res, matrix[i][j])
			j--
		}
		j++
		i--
		rows--
		// left -> up
		for k = 0; k < rows; k++ {
			res = append(res, matrix[i][j])
			i--
		}
		i++
		j++
		cols--
	}

	return res[:m*n]
}
