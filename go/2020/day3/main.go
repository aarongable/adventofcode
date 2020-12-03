package main

import (
	"fmt"
	"os"

	"github.com/aarongable/adventofcode/2020/lib"
)

func part1(m [][]bool) int {
	res := 0
	width := len(m[0])
	for i, row := range m {
		if row[(i*3)%width] {
			res = res + 1
		}
	}
	return res
}

func part2(m [][]bool) int {
	return 0
}

func main() {
	part, scanner := lib.Boilerplate()
	treemap := make([][]bool, 0)
	for scanner.Scan() {
		row := make([]bool, 0)
		for _, c := range scanner.Text() {
			if c == '.' {
				row = append(row, false)
			} else if c == '#' {
				row = append(row, true)
			} else {
				fmt.Printf("Encountered unexpected character: %s\n", c)
				os.Exit(1)
			}
		}
		treemap = append(treemap, row)
	}
	if part == 1 {
		fmt.Println(part1(treemap))
	} else {
		fmt.Println(part2(treemap))
	}
}
