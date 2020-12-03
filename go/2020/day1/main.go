package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aarongable/adventofcode/2020/lib"
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
	part, scanner := lib.Boilerplate()
	nums := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Failed to convert input line to int: %s\n", scanner.Text())
			os.Exit(1)
		}
		nums = append(nums, i)
	}
	if part == 1 {
		fmt.Println(part1(nums))
	} else {
		fmt.Println(part2(nums))
	}
}
