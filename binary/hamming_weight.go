package binary

import "fmt"

func HammingWeightMain() {
	fmt.Println(hammingWeight(1))
	fmt.Println(hammingWeight(8))
	fmt.Println(hammingWeight(7))
	fmt.Println(hammingWeight(0b1111111111111111 << 5))
}

func hammingWeight(num uint32) int {
	var cnt uint32 = 0
	for num != 0 {
		cnt += num & 1
		num >>= 1
	}

	return int(cnt)
}
