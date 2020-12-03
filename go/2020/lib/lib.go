package lib

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Boilerplate performs setup that is expected to be common to all AoC
// problems: it declares flags to differentiate the Part 1 / Part 2
// solutions and to locate the input file. It returns an int representing
// which part of the problem should be solved, and a Scanner over the lines
// of the opened input file for the solution to parse as appropriate.
func Boilerplate() (int, *bufio.Scanner) {
	part := flag.Int("part", 1, "Solve part one or part two of the problem")
	flag.Parse()
	if *part < 1 || *part > 2 {
		fmt.Println("Please select part 1 or part 2")
		os.Exit(1)
	}
	if len(flag.Args()) != 1 {
		fmt.Println("Must give a path to a file containing input")
		os.Exit(1)
	}
	filename := flag.Args()[0]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to open file")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	return *part, scanner
}
