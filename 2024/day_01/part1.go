package day_01

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mnezh/advent-of-code/2024/util"
)

func Do() {
	fmt.Println(Calc("day_01/input"))
}

func Calc(filePath string) int {
	return distance(readInput(util.DataPath(filePath)))
}

func distance(left, right []int) int {
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
	slices.Sort(left)
	slices.Sort(right)
	return left, right
}
