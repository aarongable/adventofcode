package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/aarongable/adventofcode/2020/lib"
)

type policy struct {
	char  string
	lower int
	upper int
}

var policyRegexp = regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)

func parseLine(l string) (*line, error) {
	m := policyRegexp.FindStringSubmatch(l)
	lower, err := strconv.Atoi(m[1])
	if err != nil {
		return nil, err
	}
	upper, err := strconv.Atoi(m[2])
	if err != nil {
		return nil, err
	}
	return &line{
		rule: policy{
			char:  m[3],
			lower: lower,
			upper: upper,
		},
		password: m[4],
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
		first := string(l.password[l.rule.lower-1]) == l.rule.char
		second := string(l.password[l.rule.upper-1]) == l.rule.char
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
		l, err := parseLine(scanner.Text())
		if err != nil {
			fmt.Printf("Failed to parse input line: %s\n", scanner.Text())
			os.Exit(1)
		}
		lines = append(lines, *l)
	}
	if part == 1 {
		fmt.Println(part1(lines))
	} else {
		fmt.Println(part2(lines))
	}
}
