package graphs

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/tour/tree"
)

func TreeMain() {
	rand.Seed(time.Now().UnixMilli())
	t1 := tree.New(4)
	t2 := tree.New(3)

	fmt.Println(Same(t1, t2))
}

func Same(t1, t2 *tree.Tree) bool {
	pairs := 0
	ch := make(chan int)

	go func() {
		defer close(ch)
		Walk(t1, ch)
		Walk(t2, ch)
	}()

	for v := range ch {
		pairs ^= v
	}

	return pairs == 0
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}
