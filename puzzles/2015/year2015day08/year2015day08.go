package year2015day08

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/puzzles/registry"
)

//go:embed year2015day08.txt
var input string

func init() {
	// https://adventofcode.com/2015/day/8
	registry.AddPuzzle(2015, 8, "Matchsticks", input, part1, part2)
}

func part1(input string) string {
	return strconv.Itoa(0)
}

func part2(input string) string {
	return strconv.Itoa(0)
}
