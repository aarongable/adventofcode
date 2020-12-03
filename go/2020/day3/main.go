package main

import (
	"fmt"
	"os"

	"github.com/aarongable/adventofcode/2020/lib"
)

func countTrees(m [][]bool, slopeX int, slopeY int) int {
	res := 0
	width := len(m[0])
	for x, y := 0, 0; y < len(m); x, y = x+slopeX, y+slopeY {
		if m[y][x%width] {
			res = res + 1
		}
	}
	return res
}

func part1(m [][]bool) int {
	return countTrees(m, 3, 1)
}

func part2(m [][]bool) int {
	a := countTrees(m, 1, 1)
	b := countTrees(m, 3, 1)
	c := countTrees(m, 5, 1)
	d := countTrees(m, 7, 1)
	e := countTrees(m, 1, 2)
	t := a * b * c * d * e
	return t
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
