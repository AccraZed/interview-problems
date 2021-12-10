package dp

import "fmt"

func FibMain() {
	memo := make(map[int]int, 1)
	memo[0] = 0
	memo[1] = 1
	fmt.Println(memoFib(50, memo))
	fmt.Println(dpFib(50))
}

func memoFib(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}

	val := memoFib(n-1, memo)
	memo[n-1] = val

	return val + memoFib(n-2, memo)
}

func dpFib(n int) int {
	vals := make([]int, n+1)

	vals[0] = 0
	vals[1] = 1

	for i := 2; i <= n; i++ {
		vals[i] = vals[i-1] + vals[i-2]
	}

	return vals[n]
}
