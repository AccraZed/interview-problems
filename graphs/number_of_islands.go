package graphs

func NumberOfIslandsMain() {

}

func numberOfIslands(board [][]int) int {
	cnt := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 1 {
				cnt++
				sinkIsland(board, i, j)
			}
		}
	}

	return cnt
}

func sinkIsland(board [][]int, i, j int) {
	type coord struct {
		i int
		j int
	}
	stack := make([]coord, 0)
	stack = append(stack, coord{i, j})
	board[i][j] = 0

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, dir := range dirs {
			if board[top.i+dir[0]][top.j+dir[1]] == 1 {
				board[top.i+dir[0]][top.j+dir[1]] = 0
				stack = append(stack, coord{i: top.i + dir[0], j: top.j + dir[1]})
			}
		}
	}
}
