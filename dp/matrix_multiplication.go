package dp

import (
	"fmt"
	"math"
)

// Find the minimum number of mult operations needed to multiply X matrices together
// Assume the matrices can be multiplied

func btCntMults(d []int, i, j int) int {
	if i == j {
		return 0
	}

	val := math.MaxInt
	for k := i; k < j; k++ {
		val = min(val, btCntMults(d, i, k)+btCntMults(d, k+1, j)+(d[i]*d[k+1]*d[j+1]))
	}

	return val
}

func memoCntMults(d []int, i, j int, memo [][]int) int {
	if i == j {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	val := math.MaxInt
	for k := i; k < j; k++ {
		val = min(val, btCntMults(d, i, k)+btCntMults(d, k+1, j)+(d[i]*d[k+1]*d[j+1]))
	}

	memo[i][j] = val

	return val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dpCountMults(d []int) int {
	n := len(d) - 1
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
		res[i][i] = 0
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			val := math.MaxInt
			for k := i; k < j; k++ {
				val = min(val, res[i][k]+res[k+1][j]+d[i]*d[k+1]*d[j+1])
			}

			res[i][j] = val
		}
	}

	return res[0][n-1]
}

func MatrixMultMain() {
	d := []int{2, 4, 2, 3, 1}

	memo := make([][]int, len(d)-1)

	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(d)-1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}

	fmt.Println(btCntMults(d, 0, len(d)-2))
	fmt.Println(memoCntMults(d, 0, len(d)-2, memo))
	fmt.Println(dpCountMults(d))
}
