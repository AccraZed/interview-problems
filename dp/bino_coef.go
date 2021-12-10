package dp

import "fmt"

type Key struct {
	n, k int
}

func BinoMain() {
	memo := make(map[Key]int)
	fmt.Println(betterBinoCoef(40, 20, memo))
	fmt.Println(binoCoef(40, 20))
}

func betterBinoCoef(n, k int, memo map[Key]int) int {
	if k == n || k == 0 {
		return 1
	}
	if k > n || k < 0 {
		return 0
	}

	var first int
	if chose, ok := memo[Key{n: n - 1, k: k - 1}]; ok {
		first = chose
	} else {
		first = betterBinoCoef(n-1, k-1, memo)
		memo[Key{n: n - 1, k: k - 1}] = first
	}

	var second int
	if chose, ok := memo[Key{n: n - 1, k: k}]; ok {
		second = chose
	} else {
		second = betterBinoCoef(n-1, k, memo)
		memo[Key{n: n - 1, k: k}] = second
	}

	return first + second
}

func binoCoef(n, k int) int {
	if k == n || k == 0 {
		return 1
	}
	if k > n || k < 0 {
		return 0
	}

	return binoCoef(n-1, k-1) + binoCoef(n-1, k)
}
