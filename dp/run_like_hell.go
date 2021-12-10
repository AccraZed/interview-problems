package dp

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func RunLikeHellMain() {
	raw, _ := os.ReadFile("25000000Blocks.txt")
	str := string(raw)
	sstr := strings.Split(str, " ")

	blocks := make([]int, 25000000)
	for i, s := range sstr {
		blocks[i], _ = strconv.Atoi(s)
	}

	var total int64 = 0
	n := 100
	for i := 0; i < n; i++ {
		start := time.Now()
		dpMaxGain(blocks)
		elapsed := time.Since(start).Milliseconds()

		total += elapsed
	}

	fmt.Println(total / int64(n))
}

func dpMaxGain(blocks []int) int {
	res := make([]int, 2)
	if len(blocks) != 0 {
		res[1] = blocks[0]
	}

	for k := 2; k <= len(blocks); k++ {
		res[k%2] = max(res[(k-2)%2]+blocks[k-1], res[(k-1)%2])
	}

	return res[len(blocks)%2]
}
