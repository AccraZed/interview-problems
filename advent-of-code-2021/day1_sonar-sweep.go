package advent2021

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Day1() {
	f, err := os.Open(filepath.Join("advent-of-code-2021", "day1_input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		depth, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, depth)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d1part1(data))
	fmt.Println(d1part2(data))
}

func d1part1(data []int) int {
	cnt := 0

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			cnt++
		}
	}

	return cnt
}

func d1part2(data []int) int {
	cnt := 0

	for i := 3; i < len(data); i++ {
		if data[i] > data[i-3] {
			cnt++
		}
	}

	return cnt
}
