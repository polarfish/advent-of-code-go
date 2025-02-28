package puzzles

import (
	_ "embed"
	"strconv"
)

//go:embed year2015day01.txt
var year2015Day01Input string

func init() {
	addPuzzle(2015, 1, "Not Quite Lisp", year2015Day01Input, year2015Day01Part1, year2015Day01Part2)
}

func year2015Day01Part1(input string) string {
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

func year2015Day01Part2(input string) string {
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
	return errorResult
}
