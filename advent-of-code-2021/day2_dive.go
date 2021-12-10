package advent2021

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Day2() {
	hori, depth := d2part1()
	fmt.Println(hori, depth)
	fmt.Println(hori * depth)

	hori, depth = d2part2()
	fmt.Println(hori, depth)
	fmt.Println(hori * depth)
}

func d2part1() (hori int, depth int) {
	f, err := os.Open(filepath.Join("advent-of-code-2021", "day2_input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	in := bufio.NewScanner(f)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		switch in.Text() {
		case "forward":
			in.Scan()
			amt, err := strconv.Atoi(in.Text())
			if err != nil {
				log.Fatal(err)
			}
			hori += amt
		case "down":
			in.Scan()
			amt, err := strconv.Atoi(in.Text())
			if err != nil {
				log.Fatal(err)
			}
			depth += amt
		case "up":
			in.Scan()
			amt, err := strconv.Atoi(in.Text())
			if err != nil {
				log.Fatal(err)
			}
			depth -= amt
		}
	}
	if err := in.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func d2part2() (hori int, depth int) {
	f, err := os.Open(filepath.Join("advent-of-code-2021", "day2_input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	in := bufio.NewScanner(f)
	in.Split(bufio.ScanWords)

	aim := 0

	for in.Scan() {
		switch in.Text() {
		case "forward":
			in.Scan()
			amt, err := strconv.Atoi(in.Text())
			if err != nil {
				log.Fatal(err)
			}
			hori += amt
			depth += aim * amt
		case "down":
			in.Scan()
			amt, err := strconv.Atoi(in.Text())
			if err != nil {
				log.Fatal(err)
			}
			aim += amt
		case "up":
			in.Scan()
			amt, err := strconv.Atoi(in.Text())
			if err != nil {
				log.Fatal(err)
			}
			aim -= amt
		}
	}
	if err := in.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
