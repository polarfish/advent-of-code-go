package year2015day01

import (
	_ "embed"

	"github.com/polarfish/advent-of-code-go/tools/registry"
	"github.com/polarfish/advent-of-code-go/tools/utils"
)

//go:embed year2015day01.txt
var input string

func init() {
	// https://adventofcode.com/2015/day/1
	registry.AddSolution(2015, 1, "Not Quite Lisp", input, part1, part2)
}

func part1(input string) string {
	floor := 0
	for _, ch := range input {
		if ch == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return utils.ToStr(floor)
}

func part2(input string) string {
	floor := 0
	for i, ch := range input {
		if ch == '(' {
			floor += 1
		} else {
			floor -= 1
		}

		if floor < 0 {
			return utils.ToStr(i + 1)
		}
	}
	return utils.ErrorResult
}
