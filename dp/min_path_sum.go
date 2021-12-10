package dp

import (
	"fmt"
	"math/rand"
	"time"
)

func MinPathSumMain() {
	m, n := 100, 100
	matrix := make([][]int, m)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = r.Int()%(m*n) + 1
		}
		fmt.Println(matrix[i])
	}

	fmt.Println(dpMinPathSum(matrix))
}

func dpMinPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	dp[0][0] = matrix[0][0]

	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}

	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = dp[0][j-1] + matrix[0][j]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[len(matrix)-1][len(matrix[0])-1]
}

func recMinPathSum(matrix [][]int, targRow, targCol int) int {
	if targCol == 1 && targRow == 1 {
		return matrix[0][0]
	}

	if targRow == 1 {
		return matrix[targRow-1][targCol-1] + recMinPathSum(matrix, targRow, targCol-1)
	}

	if targCol == 1 {
		return matrix[targRow-1][targCol-1] + recMinPathSum(matrix, targRow-1, targCol)
	}

	fromTop := recMinPathSum(matrix, targRow-1, targCol)
	fromLeft := recMinPathSum(matrix, targRow, targCol-1)
	min := min(fromTop, fromLeft)

	return matrix[targRow-1][targCol-1] + min
}
