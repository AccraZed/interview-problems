package binary

import "fmt"

// https://leetcode.com/problems/missing-number/
func MissingNumberMain() {
	fmt.Println(missingNumber([]int{3, 0, 1}))    // 2
	fmt.Println(bitMissingNumber([]int{3, 0, 1})) // 2
}

// XOR every number i from [1,n] into a variable, as well as every number num in
// nums. The number i that never had a pair will be left uncancelled
func bitMissingNumber(nums []int) int {
	res := 0

	for i, num := range nums {
		res = res ^ (i + 1) ^ num
	}

	return res
}

// Insert all values into a hashmap, iterate until you find the value that was
// never entered
func missingNumber(nums []int) int {
	seen := make(map[int]bool)

	for _, num := range nums {
		seen[num] = true
	}

	for i := 0; i <= len(nums); i++ {
		if !seen[i] {
			return i
		}
	}

	return -1
}
