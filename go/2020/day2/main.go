package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aarongable/adventofcode/2020/lib"
)

type policy struct {
	char  string
	lower int
	upper int
}

func parsePolicy(p string) (*policy, error) {
	s := strings.Split(p, " ")
	limits := strings.Split(s[0], "-")
	lower, err := strconv.Atoi(limits[0])
	if err != nil {
		return nil, err
	}
	upper, err := strconv.Atoi(limits[1])
	if err != nil {
		return nil, err
	}
	return &policy{
		char:  s[1],
		lower: lower,
		upper: upper,
	}, nil
}

type line struct {
	rule     policy
	password string
}

func part1(lines []line) int {
	res := 0
	for _, l := range lines {
		c := strings.Count(l.password, l.rule.char)
		if l.rule.lower <= c && c <= l.rule.upper {
			res = res + 1
		}
	}
	return res
}

func part2(lines []line) int {
	res := 0
	for _, l := range lines {
		first := string(l.password[l.rule.lower]) == l.rule.char
		second := string(l.password[l.rule.upper]) == l.rule.char
		if first != second {
			res = res + 1
		}
	}
	return res
}

func main() {
	part, scanner := lib.Boilerplate()
	lines := make([]line, 0)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")
		p, err := parsePolicy(split[0])
		if err != nil {
			fmt.Printf("Failed to parse input line: %s\n", scanner.Text())
			os.Exit(1)
		}
		lines = append(lines, line{rule: *p, password: split[1]})
	}
	if part == 1 {
		fmt.Println(part1(lines))
	} else {
		fmt.Println(part2(lines))
	}
}
