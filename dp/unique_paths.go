package dp

import "fmt"

func backtrackUniquePaths(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	return backtrackUniquePaths(m-1, n) + backtrackUniquePaths(m, n-1)
}

func dpUniquePaths(m, n int) int {
	res := make([][]int, m)

	// Init array, fill base cases
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
		res[i][0] = 1
	}
	for i := 0; i < len(res[0]); i++ {
		res[0][i] = 1
	}

	// Do the DP
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}

func UniquePathsMain() {
	fmt.Println("DP Func:")
	fmt.Println(dpUniquePaths(2, 2))
	fmt.Println(dpUniquePaths(2, 3))
	fmt.Println(dpUniquePaths(3, 3))
	fmt.Println()

	fmt.Println("Backtrack Func:")
	fmt.Println(backtrackUniquePaths(2, 2))
	fmt.Println(backtrackUniquePaths(2, 3))
	fmt.Println(backtrackUniquePaths(3, 3))
}
