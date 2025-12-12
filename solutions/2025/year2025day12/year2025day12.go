package year2025day12

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/tools/registry"
)

//go:embed year2025day12.txt
var input string

func init() {
	// https://adventofcode.com/2025/day/12
	registry.AddSolution(2025, 12, "Christmas Tree Farm", input, part1, part2)
}

func part1(input string) (string, error) {
	return strconv.Itoa(0), nil
}

func part2(input string) (string, error) {
	return strconv.Itoa(0), nil
}
