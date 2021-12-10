package dp

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/combination-sum-iv/

func CombinationSumMain() {
	fmt.Println(recCombinationSum([]int{1, 2, 3}, 4))      // 7
	fmt.Println(dpCombinationSum([]int{1, 2, 3}, 4))       // 7
	fmt.Println(dpGreedyCombinationSum([]int{1, 2, 3}, 4)) // 7
}

func dpGreedyCombinationSum(nums []int, target int) int {
	sort.Ints(nums)
	res := make([]int, target+1)
	res[0] = 1

	for curTarg := 1; curTarg <= target; curTarg++ {
		for _, num := range nums {
			if num > curTarg {
				break
			}
			res[curTarg] += res[curTarg-num]
		}
	}

	return res[target]
}

func dpCombinationSum(nums []int, target int) int {
	res := make([]int, target+1)
	res[0] = 1

	for curTarg := 1; curTarg <= target; curTarg++ {
		for _, num := range nums {
			if num <= curTarg {
				res[curTarg] += res[curTarg-num]
			}
		}
	}

	return res[target]
}

func recCombinationSum(nums []int, target int) int {
	if target == 0 {
		return 1
	}

	cnt := 0
	for _, num := range nums {
		if num <= target {
			cnt += recCombinationSum(nums, target-num)
		}
	}

	return cnt
}
