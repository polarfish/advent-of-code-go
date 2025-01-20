package year2015day01

import (
	_ "embed"
	"strconv"

	"github.com/polarfish/advent-of-code-go/utils"
)

//go:embed day01.txt
var input string

func New() *utils.Puzzle {
	return &utils.Puzzle{
		Year:  2015,
		Day:   1,
		Name:  "Not Quite Lisp",
		Input: input,
		Part1: Part1,
		Part2: Part2,
	}
}

func Part1(input string) string {
	floor := 0
	for _, ch := range input {
		if ch == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return strconv.Itoa(floor)
}

func Part2(input string) string {
	floor := 0
	for i, ch := range input {
		if ch == '(' {
			floor += 1
		} else {
			floor -= 1
		}

		if floor < 0 {
			return strconv.Itoa(i + 1)
		}
	}
	return utils.ERR
}
