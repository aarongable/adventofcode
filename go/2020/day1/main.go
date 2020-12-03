package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func part1(nums []int) int {
	l := len(nums)
	for i, x := range nums[0 : l-1] {
		for _, y := range nums[i:l] {
			if x+y == 2020 {
				return x * y
			}
		}
	}
	return 0
}

func part2(nums []int) int {
	l := len(nums)
	for i, x := range nums[0 : l-2] {
		for j, y := range nums[i : l-1] {
			for _, z := range nums[j:l] {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return 0
}

func main() {
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
	nums := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Failed to convert input line to int: %s\n", scanner.Text())
			os.Exit(1)
		}
		nums = append(nums, i)
	}
	if *part == 1 {
		fmt.Println(part1(nums))
	} else {
		fmt.Println(part2(nums))
	}
}
