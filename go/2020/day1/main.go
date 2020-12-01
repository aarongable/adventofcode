package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
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
	for i, x := range nums[0 : len(nums)-1] {
		for _, y := range nums[i:len(nums)] {
			if x+y == 2020 {
				fmt.Printf("%d * %d = %d\n", x, y, x*y)
				os.Exit(0)
			}
		}
	}
}
