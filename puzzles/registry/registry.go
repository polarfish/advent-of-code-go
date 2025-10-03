package registry

import (
	"github.com/polarfish/advent-of-code-go/puzzles/utils"
)

var allPuzzles []*utils.Puzzle

func addPuzzle(year int, day int, name string, input string, part1 func(string) string, part2 func(string) string) {
	allPuzzles = append(allPuzzles, &utils.Puzzle{
		Year:  year,
		Day:   day,
		Name:  name,
		Input: input,
		Part1: part1,
		Part2: part2,
	})
}

func AddPuzzle(year int, day int, name string, input string, part1 func(string) string, part2 func(string) string) {
	addPuzzle(year, day, name, input, part1, part2)
}

func GetAllPuzzles() []*utils.Puzzle {
	return allPuzzles
}
