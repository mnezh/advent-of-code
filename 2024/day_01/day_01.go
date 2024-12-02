package day_01

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mnezh/advent-of-code/2024/util"
)

func Day1(part int) int {
	if part == 1 {
		return Calc("day_01/input", distance)
	}
	return Calc("day_01/input", similarityScore)
}

func distance(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for index, l := range left {
		diff := l - right[index]
		if diff < 0 {
			diff = -diff
		}
		sum = sum + diff
	}
	return sum
}

func similarityScore(left, right []int) int {
	score := 0
	for _, l := range left {
		score += l * occurrences(right, l)
	}
	return score
}

func occurrences(list []int, lookingFor int) int {
	res := 0
	for _, item := range list {
		if item == lookingFor {
			res++
		}
	}
	return res
}

func readInput(filePath string) ([]int, []int) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var left, right []int
	for _, line := range strings.Split(string(content), "\n") {
		if len(line) > 0 {
			fields := strings.Split(line, "   ")
			l, _ := strconv.Atoi(fields[0])
			r, _ := strconv.Atoi(fields[1])
			left = append(left, l)
			right = append(right, r)
		}
	}

	return left, right
}

func Calc(filePath string, calculator func([]int, []int) int) int {
	return calculator(readInput(util.DataPath(filePath)))
}
