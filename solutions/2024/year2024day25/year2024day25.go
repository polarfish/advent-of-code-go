package year2024day25

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/registry"
)

//go:embed year2024day25.txt
var input string

func init() {
	// https://adventofcode.com/2024/day/25
	registry.AddSolution(2024, 25, "Code Chronicle", input, part1, part2)
}

func part1(input string) string {
	return strconv.Itoa(0)
}

func part2(input string) string {
	return strconv.Itoa(0)
}
