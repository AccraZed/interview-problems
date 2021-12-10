package dp

import "fmt"

func JumpMain() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(minJumps(nums, 0))
}

func minJumps(nums []int, k int) int {
	if k >= len(nums)-1 {
		return 0
	}
	if nums[k] == 0 {
		return -1
	}

	val := 0
	for i := 1; i <= nums[k]; i++ {
		curJumps := minJumps(nums, k+i)
		if curJumps == -1 {
			continue
		}

		val = min(val, curJumps)
	}

	return val + 1
}
