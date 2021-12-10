package dp

import (
	"fmt"
)

type item struct {
	w int
	v int
}

type total struct {
	tW int
	tV int
}

// time complexity: O(2^n)
// space complexity: O(1)
func backtrackKnapsack(items []item, cap, k int) int {
	if k == -1 || cap == 0 {
		return 0
	}

	var val int
	if items[k].w <= cap {
		val = max(
			backtrackKnapsack(items, cap-items[k].w, k-1)+items[k].v,
			backtrackKnapsack(items, cap, k-1),
		)
	} else {
		val = backtrackKnapsack(items, cap, k-1)
	}

	return val
}

func dpKnapsack(items []item, cap int) int {
	res := make([][]int, len(items)+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, cap+1)
	}

	for iI := 1; iI <= len(items); iI++ {
		for iW := 1; iW <= cap; iW++ {
			if items[iI-1].w <= iW {
				res[iI][iW] = max(
					res[iI-1][iW-items[iI-1].w]+items[iI-1].v,
					res[iI-1][iW],
				)
			} else {
				res[iI][iW] = res[iI-1][iW]
			}
		}
	}

	return res[len(items)][cap]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func KnapsackMain() {
	items := []item{
		{w: 4, v: 6},
		{w: 2, v: 4},
		{w: 3, v: 5},
		{w: 1, v: 3},
		{w: 6, v: 9},
		{w: 4, v: 7},
	}

	used := make(map[int]bool)
	for i, _ := range items {
		used[i] = false
	}

	maxWeight := 10
	fmt.Println(dpKnapsack(items, maxWeight))
	fmt.Println(backtrackKnapsack(items, maxWeight, len(items)-1))
}
