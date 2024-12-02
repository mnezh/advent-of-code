package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mnezh/advent-of-code/2024/day_01"
)

func intArg(num int) int {
	arg, err := strconv.Atoi(os.Args[num])
	if err != nil {
		fmt.Printf("Not an integer: %s\n", os.Args[num])
		os.Exit(2)
	}
	return arg
}

var parts = map[int]func(int) int{
	1: day_01.Day1,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aoc <day> [part]")
		os.Exit(1)
	}

	day := intArg(1)
	part := 1
	if len(os.Args) > 2 {
		part = intArg(2)
	}
	fmt.Println(parts[day](part))
}
