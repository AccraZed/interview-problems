package dp

import "fmt"

func CntUniqueBTMain() {
	for i := 1; i < 10; i++ {
		fmt.Printf("F(%d): %d\n", i, cntUniqueBTs(i))
	}
}

func cntUniqueBTs(num int) int {
	res := make([]int, num+1)
	res[0] = 1

	for n := 1; n <= num; n++ {
		for i := 1; i <= n; i++ {
			res[n] += res[i-1] * res[n-i]
		}
	}

	return res[num]
}
