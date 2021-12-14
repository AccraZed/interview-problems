package advent2021

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Day3() {
	// f, err := os.Open(filepath.Join("advent-of-code-2021", "day3_input.txt"))
	f, err := os.Open(filepath.Join("advent-of-code-2021", "day3_input2.txt"))
	if err != nil {
		log.Fatal(err)
	}

	input := make([]uint64, 0, 1000)

	in := bufio.NewScanner(f)
	for in.Scan() {
		v, err := strconv.ParseUint(in.Text(), 2, len(in.Text()))
		if err != nil {
			log.Fatal(err)
		}

		input = append(input, v)
	}

	fmt.Println(day3part1(input, 5))
	fmt.Println(day3part2(input, 5))
}

func day3part1(input []uint64, bitlen int) uint64 {
	freqs := make([]int, bitlen)

	for pos := 0; pos < bitlen; pos++ {
		switch mostFreq(input, pos, bitlen) {
		}
	}

	var gamma uint64 = 0
	for pos := len(freqs) - 1; pos >= 0; pos-- {
		gamma <<= 1
		if freqs[pos] > 0 {
			gamma++
		}
	}
	return gamma * (^gamma & (1<<bitlen - 1))
}

func day3part2(input []uint64, bitlen int) int {

	fmt.Printf("%b\n", input)

	return 0
}

func mostFreq(input []uint64, pos, bitlen int) int {
	balance := 0
	for _, num := range input {
		if (1<<pos)&input[num] != 0 {
			balance++
		} else {
			balance--
		}
	}

	if balance > 0 {
		return 1
	} else if balance < 0 {
		return -1
	} else {
		return 0
	}
}
