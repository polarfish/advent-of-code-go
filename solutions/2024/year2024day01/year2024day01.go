package year2024day01

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2024day01.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/1
	registry.AddSolution(2024, 1, "Historian Hysteria", input, part1, part2)
}

func part1(input string) (string, error) {
	var err error
	lines := utils.Lines(input)
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		split := strings.Split(line, "   ")
		left[i], err = strconv.Atoi(split[0])
		if err != nil {
			return "", err
		}
		right[i], err = strconv.Atoi(split[1])
		if err != nil {
			return "", err
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	var total int
	for i := 0; i < len(lines); i++ {
		total += abs(left[i] - right[i])
	}

	return strconv.Itoa(total), nil
}

func part2(input string) (string, error) {
	lines := utils.Lines(input)
	left := make([]int, len(lines))
	right := make(map[int]int, len(lines))
	for i, line := range lines {
		split := strings.Split(line, "   ")
		l, err := strconv.Atoi(split[0])
		if err != nil {
			return "", err
		}
		r, err := strconv.Atoi(split[1])
		if err != nil {
			return "", err
		}
		left[i] = l
		right[r]++
	}

	var total int
	for _, n := range left {
		if val, exists := right[n]; exists {
			total += n * val
		}
	}

	return strconv.Itoa(total), nil
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
