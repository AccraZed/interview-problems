package dp

import (
	"fmt"
	"math"
)

func SubseqMain() {
	fmt.Println(lcss([]rune("goodmorning"), []rune("hodor")))
	fmt.Println(lcss([]rune("cat"), []rune("cant")))
	fmt.Println(lcss([]rune("hodor"), []rune("goodmorning")))
	fmt.Println(lcss([]rune("goodmorning"), []rune("hodor")))
}

func lcss(a, b []rune) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	if a[len(a)-1] == b[len(b)-1] {
		return 1 + lcss(a[:len(a)-1], b[:len(b)-1])
	}

	return int(math.Max(float64(lcss(a[:len(a)-1], b)), float64(lcss(a, b[:len(b)-1]))))
}
