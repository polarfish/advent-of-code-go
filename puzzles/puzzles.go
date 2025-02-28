package puzzles

import (
	"time"
)

const errorResult string = "error"
const notApplicableResult string = "n/a"

type Puzzle struct {
	Day   int
	Year  int
	Name  string
	Input string
	Part1 func(input string) string
	Part2 func(input string) string
}

type Result struct {
	Puzzle    *Puzzle
	Result1   string
	Result2   string
	Duration1 time.Duration
	Duration2 time.Duration
}

var allPuzzles []*Puzzle

func addPuzzle(year int, day int, name string, input string, part1 func(string) string, part2 func(string) string) {
	allPuzzles = append(allPuzzles, &Puzzle{
		Year:  year,
		Day:   day,
		Name:  name,
		Input: input,
		Part1: part1,
		Part2: part2,
	})
}

func GetAllPuzzles() []*Puzzle {
	return allPuzzles
}
